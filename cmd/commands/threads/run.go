package threads

import (
	"fmt"
	"openAIThreader/internal/data"
	"openAIThreader/internal/openai"

	"github.com/spf13/cobra"
)

// NewRunCommand initializes the 'run' command.
func NewRunCommand(client openai.Client, frame *data.Frame) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Execute the currently selected Thread and stream responses",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := frame.HasSelection(); err != nil {
				return err
			}

			err := client.RunThread(frame.ThreadID, frame.AssistantID)
			if err != nil {
				return fmt.Errorf("failed to run Thread: %w", err)
			}

			fmt.Println("Thread execution complete, please run fetch to view the results")

			return nil
		},
	}

	return cmd
}
