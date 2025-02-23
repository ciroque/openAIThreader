package openai

import (
	"fmt"
	"slices"
)

const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
)

func RenderMessages(messages []Message) {
	slices.Reverse(messages)

	for _, message := range messages {
		role := message.Role
		color := Green
		if role == "assistant" {
			color = Cyan
		}

		for _, content := range message.Content {
			fmt.Printf("%s\n%s:\n %s\n\n", colorize(message.ID, Yellow), bold(colorize(role, color)), content.Text.Value)
		}
	}
}

// Format text with bold
func bold(text string) string {
	return Bold + text + Reset
}

// Format text with color
func colorize(text, color string) string {
	return color + text + Reset
}
