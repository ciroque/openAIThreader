package threads

import (
	"fmt"
	"openAIThreader/internal/data"
	"openAIThreader/internal/storage"

	"github.com/spf13/cobra"
)

func NewSelectCommand(store storage.Provider, frame *data.Frame) *cobra.Command {
	var threadID string

	cmd := &cobra.Command{
		Use:   "select",
		Short: "Select a Thread",
		RunE: func(cmd *cobra.Command, args []string) error {
			threads, err := store.LoadThreads()
			if err != nil {
				return fmt.Errorf("failed to load Threads: %w", err)
			}

			name, found := findByValue(threads, threadID)
			if !found {
				return fmt.Errorf("unable to retrieve name for Thread: %s", threadID)
			}

			frame.Select(threadID, name)

			fmt.Printf("Selected Thread %s (ID: %s)\n", name, threadID)
			return nil
		},
	}

	cmd.Flags().StringVarP(&threadID, "threadId", "t", "user", "The ID of the Thread to select")
	cmd.MarkFlagRequired("threadId")

	return cmd
}

func findByValue(m map[string]string, val string) (string, bool) {
	for k, v := range m {
		if v == val {
			return k, true
		}
	}
	return "", false
}
