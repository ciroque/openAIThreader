package threads

import (
	"fmt"
	"openAIThreader/internal/data"
	"openAIThreader/internal/openai"
	"openAIThreader/internal/storage"

	"github.com/spf13/cobra"
)

func NewDeleteCommend(client openai.Client, store storage.Provider, frame *data.Frame) *cobra.Command {
	var threadId string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a Thread",
		RunE: func(cmd *cobra.Command, args []string) error {
			threadID, err := cmd.Flags().GetString("threadId")
			if err != nil {
				return err
			}

			if err = client.DeleteThread(threadID); err != nil {
				return err
			}

			if err := store.DeleteThread(threadID); err != nil {
				return err
			}

			frame.Clear()

			fmt.Printf("Deleted Thread %s\n", threadID)

			return nil
		},
	}

	cmd.Flags().StringVarP(&threadId, "threadId", "t", "user", "The ID of the Thread to delete")
	cmd.MarkFlagRequired("threadId")

	return cmd
}
