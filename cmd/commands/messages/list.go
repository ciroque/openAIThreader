package messages

import (
	"openAIThreader/internal/data"
	"openAIThreader/internal/openai"

	"github.com/spf13/cobra"
)

func NewListCommand(_ openai.Client, frame *data.Frame) *cobra.Command {
	var role string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists messages on the current thread",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := frame.IsLoaded(); err != nil {
				return err
			}

			openai.RenderMessages(frame.Thread.Data)

			return nil
		},
	}

	cmd.Flags().StringVarP(&role, "role", "r", "", "Filter by the Message Role (user, assistant)")

	return cmd
}
