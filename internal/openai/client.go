package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client defines the interface for OpenAI Threads API interactions.
type Client interface {
	AddMessage(threadID, role, content string) error
	CreateThread() (string, error)
	DeleteThread(threadID string) error
	FetchThreadMessages(threadID string) ([]byte, error)
	RunThread(threadID string, assistantID string) error
	RemoveMessage(threadID string, messageID string) error
}

// OpenAIClient is the concrete implementation of Client.
type OpenAIClient struct {
	apiKey     string
	httpClient *http.Client
	baseURL    string
}

// RemoveMessage removes a message from a thread.
func (c *OpenAIClient) RemoveMessage(threadID string, messageID string) error {
	url := fmt.Sprintf("%s/threads/%s/messages/%s", c.baseURL, threadID, messageID)

	reqBody := []byte("{}")

	req, _ := http.NewRequest("DELETE", url, bytes.NewBuffer(reqBody))
	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create thread: %s, %s", resp.Status, body)
	}

	return nil
}

// NewClient initializes a new OpenAI Threads client.
func NewClient(apiKey string, httpClient *http.Client) Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 30 * time.Second}
	}

	return &OpenAIClient{
		apiKey:     apiKey,
		httpClient: httpClient,
		baseURL:    "https://api.openai.com/v1",
	}
}

func (c *OpenAIClient) CreateThread() (string, error) {
	url := fmt.Sprintf("%s/threads", c.baseURL)

	// Empty JSON object to match OpenAI API spec
	reqBody := []byte("{}")

	// Create request
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	c.setHeaders(req)

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Log full response if request fails
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to create thread: %s", resp.Status)
	}

	// Parse response
	var res map[string]interface{}
	if err := json.Unmarshal(body, &res); err != nil {
		return "", err
	}

	// Extract thread ID
	threadID, exists := res["id"].(string)
	if !exists {
		return "", errors.New("thread ID missing in response")
	}

	return threadID, nil
}

func (c *OpenAIClient) DeleteThread(threadId string) error {
	url := fmt.Sprintf("%s/threads/%s", c.baseURL, threadId)

	// Empty JSON object to match OpenAI API spec
	reqBody := []byte("{}")

	// Create request
	req, _ := http.NewRequest("DELETE", url, bytes.NewBuffer(reqBody))
	c.setHeaders(req)

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Log full response if request fails
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create thread: %s, %s", resp.Status, body)
	}

	return nil
}

// AddMessage adds a message to an existing thread.
func (c *OpenAIClient) AddMessage(threadID, role, content string) error {
	url := fmt.Sprintf("%s/threads/%s/messages", c.baseURL, threadID)

	reqBody, _ := json.Marshal(map[string]string{
		"role":    role,
		"content": content,
	})

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse response
	var res map[string]interface{}
	if err := json.Unmarshal(body, &res); err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to add message: %s", resp.Status)
	}

	return nil
}

func (c *OpenAIClient) RunThread(threadID, assistantID string) error {
	url := fmt.Sprintf("%s/threads/%s/runs", c.baseURL, threadID)

	reqBody, _ := json.Marshal(map[string]string{
		"assistant_id": assistantID,
	})

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var res map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return err
	}

	runID, _ := res["id"].(string)
	fmt.Printf("Run Started: %s\n", runID)

	if err := c.WaitForRunCompletion(threadID, runID); err != nil {
		return err
	}

	return nil
}

// FetchThreadMessages retrieves the latest messages in the thread.
func (c *OpenAIClient) FetchThreadMessages(threadID string) ([]byte, error) {
	url := fmt.Sprintf("%s/threads/%s/messages", c.baseURL, threadID)

	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// WaitForRunCompletion polls the OpenAI API until the run completes.
func (c *OpenAIClient) WaitForRunCompletion(threadID, runID string) error {
	url := fmt.Sprintf("%s/threads/%s/runs/%s", c.baseURL, threadID, runID)

	for {
		time.Sleep(2 * time.Second) // ✅ Poll every 2 seconds

		req, _ := http.NewRequest("GET", url, nil)
		c.setHeaders(req)

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return fmt.Errorf("error checking run status: %v", err)
		}
		defer resp.Body.Close()

		var res map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return err
		}

		// Extract status
		status, _ := res["status"].(string)
		fmt.Printf("Run Status: %s\n", status)

		if status == "completed" {
			return nil // ✅ The run is done
		} else if status == "failed" {
			return fmt.Errorf("run failed")
		}
	}
}

// setHeaders sets the authorization and content headers.
func (c *OpenAIClient) setHeaders(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("openAI-Beta", "assistants=v2")
}
