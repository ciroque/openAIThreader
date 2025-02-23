package messages

import (
	"fmt"
	"openAIThreader/internal/data"
	"openAIThreader/internal/openai"

	"github.com/spf13/cobra"
)

// NewAddCommand initializes the 'add' command.
func NewAddCommand(client openai.Client, frame *data.Frame) *cobra.Command {
	var role, content string

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a message to an existing thread",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := frame.HasSelection(); err != nil {
				return err
			}

			c, err := cmd.Flags().GetString("content")
			if err != nil {
				return err
			}
			content = c

			if content == "" {
				return fmt.Errorf("message content cannot be empty")
			}

			// ✅ Fetch role flag
			r, _ := cmd.Flags().GetString("role")
			role = r

			err = client.AddMessage(frame.ThreadID, role, content)
			if err != nil {
				return fmt.Errorf("failed to add message: %w", err)
			}

			fmt.Printf("Added Message to thread %s as %s: %s\n", frame.ThreadID, role, content)

			return nil
		},
	}

	cmd.Flags().StringVarP(&role, "role", "r", "user", "Role of the message (user, assistant)")
	cmd.Flags().StringVarP(&content, "content", "c", "", "Message content (required)")
	cmd.MarkFlagRequired("content")

	return cmd
}
