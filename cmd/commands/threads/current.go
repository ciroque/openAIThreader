package threads

import (
	"fmt"
	"openAIThreader/internal/data"

	"github.com/spf13/cobra"
)

func NewCurrentCommand(frame *data.Frame) *cobra.Command {
	return &cobra.Command{
		Use:   "current",
		Short: "Show the currently selected Thread",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := frame.HasSelection(); err != nil {
				return err
			}

			fmt.Printf("Current Thread: %s (ID: %s)\n", frame.ThreadName, frame.ThreadID)

			return nil
		},
	}
}
