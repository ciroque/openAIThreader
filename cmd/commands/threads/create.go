package threads

import (
	"fmt"
	"openAIThreader/internal/data"
	"openAIThreader/internal/openai"

	"openAIThreader/internal/storage"

	"github.com/spf13/cobra"
)

// NewCreateCommand initializes the 'create' command.
func NewCreateCommand(client openai.Client, store storage.Provider, frame *data.Frame) *cobra.Command {
	var threadName string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new named Thread",
		RunE: func(cmd *cobra.Command, args []string) error {
			threadID, err := client.CreateThread()
			if err != nil {
				return fmt.Errorf("failed to create Thread: %w", err)
			}

			if err := store.SaveThreadsWithNames(threadName, threadID); err != nil {
				return fmt.Errorf("failed to save Thread: %w", err)
			}

			frame.ThreadID = threadID
			frame.ThreadName = threadName
			frame.Thread = nil

			fmt.Printf("Created Thread %s (ID: %s)\n", threadName, threadID)

			return nil
		},
	}

	cmd.Flags().StringVarP(&threadName, "name", "n", "", "Name of the Thread (required)")
	cmd.MarkFlagRequired("name")

	return cmd
}
