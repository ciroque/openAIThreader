package threads

import (
	"fmt"
	"openAIThreader/internal/data"
	"openAIThreader/internal/openai"

	"github.com/spf13/cobra"
)

func NewFetchCommand(client openai.Client, frame *data.Frame) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "Fetch the current Thread",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := frame.HasSelection(); err != nil {
				return err
			}

			content, err := client.FetchThreadMessages(frame.ThreadID)
			if err != nil {
				return fmt.Errorf("failed to fetch thread: %w", err)
			}

			response, err := openai.UnmarshalResponse(content)
			if err != nil {
				return fmt.Errorf("failed to unmarshal response: %w", err)
			}

			frame.Thread = response

			fmt.Printf("Fetched Thread %s: (ID: %s)\n", frame.ThreadName, frame.ThreadID)

			return nil
		},
	}

	return cmd
}
