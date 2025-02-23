package threads

import (
	"fmt"
	"openAIThreader/internal/storage"

	"github.com/spf13/cobra"
)

// NewListCommand initializes the 'list' command.
func NewListCommand(store storage.Provider) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all saved Threads",
		RunE: func(cmd *cobra.Command, args []string) error {
			threads, err := store.LoadThreads()
			if err != nil {
				return fmt.Errorf("failed to load Threads: %w", err)
			}

			if len(threads) == 0 {
				return nil
			}

			for name, id := range threads {
				fmt.Printf("- %s (ID: %s)\n", name, id)
			}

			return nil
		},
	}

	return cmd
}
