package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/mattn/go-shellwords"
)

type InputHandler struct {
	reader *bufio.Reader
}

func NewInputHandler() *InputHandler {
	return &InputHandler{reader: bufio.NewReader(os.Stdin)}
}

func (ih *InputHandler) ReadInput() ([]string, error) {
	input, err := ih.reader.ReadString('\n')
	if err != nil {
		return []string{}, err
	}

	args, err := shellwords.Parse(strings.TrimSpace(input))
	if err != nil {
		return []string{}, err
	}

	return args, nil
}
