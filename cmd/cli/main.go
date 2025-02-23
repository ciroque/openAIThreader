package main

import (
	"errors"
	"fmt"
	"openAIThreader/cmd"
	"openAIThreader/internal/data"
	"openAIThreader/internal/openai"
	"openAIThreader/internal/storage"
	"os"

	"github.com/spf13/cobra"
)

type AccessTokens struct {
	apiKey      string
	assistantId string
}

func main() {

	accessTokens, err := retrieveAccessTokens()
	if err != nil {
		fmt.Println("Error retrieving access tokens:\n", err)
		os.Exit(1)
	}

	client := openai.NewClient(accessTokens.apiKey, nil) // Uses default HTTP client
	store := storage.NewStorage("threads.json")

	frame := data.Frame{
		AssistantID: accessTokens.assistantId,
		ThreadID:    "",
	}

	inputHandler := NewInputHandler()

	rootCmd := cmd.NewRootCommand(client, store, &frame)

	fmt.Println("OpenAIThreader CLI - Type 'exit' or 'quit' to exit")
	startREPL(rootCmd, inputHandler)
}

// startREPL starts a Read-Eval-Print-Loop (REPL) for interactive command execution.
func startREPL(rootCmd *cobra.Command, inputHandler *InputHandler) {
	for {
		fmt.Print("openAIThreader> ")

		args, err := inputHandler.ReadInput()
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		if len(args) == 0 {
			continue
		}

		if args[0] == "exit" || args[0] == "quit" {
			fmt.Println("Exiting OpenAIThreader CLI...")
			os.Exit(0)
		}

		rootCmd.SetArgs(args)
		if err := rootCmd.Execute(); err != nil {
			fmt.Println("")
		}
	}
}

func retrieveAccessTokens() (*AccessTokens, error) {
	var errs []error
	var accessTokens *AccessTokens

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		errs = append(errs, fmt.Errorf("OPENAI_API_KEY is not set"))
	}

	assistantId := os.Getenv("OPENAI_ASSISTANT_ID")
	if assistantId == "" {
		errs = append(errs, fmt.Errorf(" OPENAI_ASSISTANT_ID is not set"))
	}

	if len(errs) == 0 {
		accessTokens = &AccessTokens{
			apiKey:      apiKey,
			assistantId: assistantId,
		}
	}

	return accessTokens, errors.Join(errs...)
}
