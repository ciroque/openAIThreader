package threads

import (
	"fmt"
	"openAIThreader/internal/data"
	"openAIThreader/internal/openai"
	"openAIThreader/internal/storage"

	"github.com/spf13/cobra"
)

func NewStoreCommand(client openai.Client, store storage.Provider, frame *data.Frame) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "store",
		Short: "Store the current Thread",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := frame.HasSelection(); err != nil {
				return err
			}

			thread, err := client.FetchThreadMessages(frame.ThreadID)
			if err != nil {
				return fmt.Errorf("failed to fetch Thread: %w", err)
			}

			if err := store.StoreThread(frame, thread); err != nil {
				return fmt.Errorf("failed to store Thread: %w", err)
			}

			fmt.Printf("Stored Thread %s\n", frame.ThreadName)

			return nil
		},
	}

	return cmd
}
