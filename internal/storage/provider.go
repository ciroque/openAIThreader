package storage

import "openAIThreader/internal/data"

// Provider defines the interface for managing thread storage.
type Provider interface {
	SaveThreadsWithNames(name, threadID string) error
	DeleteThread(name string) error
	LoadThreads() (map[string]string, error)
	GetThread(name string) (string, error)
	StoreThread(frame *data.Frame, thread []byte) error
}
