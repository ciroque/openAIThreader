package cmd

import (
	"openAIThreader/cmd/commands/messages"
	"openAIThreader/cmd/commands/misc"
	"openAIThreader/cmd/commands/threads"
	"openAIThreader/internal/data"
	"openAIThreader/internal/openai"
	"openAIThreader/internal/storage"

	"github.com/spf13/cobra"
)

// NewRootCommand initializes the root CLI command.
func NewRootCommand(client openai.Client, store storage.Provider, frame *data.Frame) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "openAIThreader",
		Short: "CLI tool for managing OpenAI Threads API",
	}

	// Add subcommands
	rootCmd.AddCommand(
		threads.NewCreateCommand(client, store, frame),
		threads.NewCurrentCommand(frame),
		threads.NewDeleteCommend(client, store, frame),
		threads.NewFetchCommand(client, frame),
		threads.NewListCommand(store),
		threads.NewRunCommand(client, frame),
		threads.NewSelectCommand(store, frame),
		threads.NewStoreCommand(client, store, frame),
		messages.NewMessageCommand(client, frame),
		misc.NewUsageCommand(),
	)

	return rootCmd
}
