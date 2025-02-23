package storage

import (
	"encoding/json"
	"fmt"
	"openAIThreader/internal/data"
	"os"
	"path"
	"sync"
	"time"
)

// FileStorage implements Provider using a JSON file.
type FileStorage struct {
	mu       sync.Mutex
	filePath string
}

func (s *FileStorage) StoreThread(frame *data.Frame, thread []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Unmarshal the escaped JSON string
	var unescapedJSON interface{}
	if err := json.Unmarshal(thread, &unescapedJSON); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Re-encode it properly formatted
	formattedJSON, err := json.MarshalIndent(unescapedJSON, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to re-encode JSON: %w", err)
	}

	timestamp := time.Now().UnixMilli()

	filePath := fmt.Sprintf("./%s-%s-%d.json", frame.ThreadName, frame.ThreadID, timestamp)
	filePath = path.Clean(filePath)

	return os.WriteFile(filePath, formattedJSON, 0644)
}

// NewStorage initializes a new file-based storage instance.
func NewStorage(fileName string) Provider {
	return &FileStorage{
		filePath: fileName,
	}
}

// SaveThreadsWithNames maintains a list of current Threads by name.
func (s *FileStorage) SaveThreadsWithNames(name, threadID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	threads, err := s.LoadThreads()
	if err != nil {
		return err
	}

	threads[name] = threadID

	file, err := os.Create(s.filePath)
	if err != nil {
		return fmt.Errorf("failed to open storage file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(threads)
}

func (s *FileStorage) DeleteThread(threadId string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	threads, err := s.LoadThreads()
	if err != nil {
		return err
	}

	removeByValue(threads, threadId)
	_, found := threads[threadId]
	if found {
		return fmt.Errorf("failed to remove thread: %s", threadId)
	}

	file, err := os.Create(s.filePath)
	if err != nil {
		return fmt.Errorf("failed to open storage file: %w", err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(threads)
}

func removeByValue(m map[string]string, targetValue string) {
	var keyToRemove string
	var found bool

	// Find the key associated with the target value
	for k, v := range m {
		if v == targetValue {
			keyToRemove = k
			found = true
			break
		}
	}

	// If found, delete the key
	if found {
		delete(m, keyToRemove)
	}
}

// LoadThreads retrieves all saved threads.
func (s *FileStorage) LoadThreads() (map[string]string, error) {
	threads := make(map[string]string)

	// Open file without locking to check existence
	file, err := os.Open(s.filePath)
	if os.IsNotExist(err) {
		return threads, nil // Return empty map if file doesn't exist
	} else if err != nil {
		return nil, fmt.Errorf("failed to read storage file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&threads); err != nil {
		return nil, fmt.Errorf("failed to parse storage file: %w", err)
	}

	return threads, nil
}

// GetThread retrieves a thread ID by name.
func (s *FileStorage) GetThread(name string) (string, error) {
	threads, err := s.LoadThreads()
	if err != nil {
		return "", err
	}

	threadID, exists := threads[name]
	if !exists {
		return "", fmt.Errorf("thread not found: %s", name)
	}

	return threadID, nil
}
