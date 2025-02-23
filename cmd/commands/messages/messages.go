package messages

import (
	"openAIThreader/internal/data"
	"openAIThreader/internal/openai"

	"github.com/spf13/cobra"
)

func NewMessageCommand(client openai.Client, frame *data.Frame) *cobra.Command {
	messageCommand := &cobra.Command{
		Use:     "messages",
		Short:   "Message commands",
		Aliases: []string{"msgs"},
	}

	messageCommand.AddCommand(NewAddCommand(client, frame))
	messageCommand.AddCommand(NewListCommand(client, frame))
	messageCommand.AddCommand(NewRemoveCommand(client, frame))

	return messageCommand
}
