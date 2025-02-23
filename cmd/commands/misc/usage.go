package misc

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewUsageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "usage",
		Short: "Print usage information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to the OpenAI Threader CLI tool!")
			fmt.Println("This tool allows you to interact with the OpenAI Threads API.")
			fmt.Println("")
			fmt.Println("You will need to set two environment variables:")
			fmt.Println("")
			fmt.Println("  OPENAI_API_KEY: Your OpenAI API key")
			fmt.Println("  OPENAI_ASSISTANT_ID: Your OpenAI Assistant ID")
			fmt.Println("")
			fmt.Println("See: https://platform.openai.com/docs/api-reference/assistants/createAssistant for instructions on creating an Assistant and getting the ID.")
			fmt.Println("")
			fmt.Println("To get started, use the following commands:")
			fmt.Println("  openAIThreader create (creates a Thread)")
			fmt.Println("  example: `create --name the-matrix-protagonist`")
			fmt.Println("")
			fmt.Println("  openAIThreader add (adds a message to a Thread)")
			fmt.Println("  example: `add --content 'Who is the protagonist of The Matrix?'`")
			fmt.Println("  example: `add --content 'Please limit your response to seventy words or less'`")
			fmt.Println("")
			fmt.Println("  openAIThreader run")
			fmt.Println("  example: `run`")
			fmt.Println("")
			fmt.Println("The run command will execute the Thread, await completion, and display the response.")
			fmt.Println("")
			fmt.Println("You can then add more messages to the Thread and run it again.")
			fmt.Println("")
			fmt.Println("Type 'exit' or 'quit' to exit the CLI tool.")
		},
	}

	return cmd
}
