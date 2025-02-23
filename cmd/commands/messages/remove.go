package messages

import (
	"fmt"
	"openAIThreader/internal/data"
	"openAIThreader/internal/openai"

	"github.com/spf13/cobra"
)

func NewRemoveCommand(client openai.Client, frame *data.Frame) *cobra.Command {
	var messageID string

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a message",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := frame.HasSelection(); err != nil {
				return err
			}

			err := client.RemoveMessage(frame.ThreadID, messageID)
			if err != nil {
				return err
			}

			fmt.Printf("Removed Message %s, %s\n", frame.ThreadID, messageID)

			return nil
		},
	}

	cmd.Flags().StringVarP(&messageID, "messageId", "i", "", "The ID of the message to remove")
	cmd.MarkFlagRequired("messageId")

	return cmd
}
