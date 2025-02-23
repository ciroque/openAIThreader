package data

import (
	"errors"
	"openAIThreader/internal/openai"
)

const (
	ErrorNoThreadSelected = "no thread selected"
	ErrorNoThreadLoaded   = "no thread loaded, run fetch"
)

type Frame struct {
	ThreadID    string
	ThreadName  string
	AssistantID string
	Thread      *openai.Response
}

func (f *Frame) HasSelection() error {
	if f.ThreadID == "" {
		return errors.New(ErrorNoThreadSelected)
	}

	return nil
}

func (f *Frame) IsLoaded() error {
	var errs []error
	errs = append(errs, f.HasSelection())
	if f.Thread == nil {
		errs = append(errs, errors.New(ErrorNoThreadLoaded))
	}

	return errors.Join(errs...)
}

func (f *Frame) Clear() {
	f.ThreadID = ""
	f.ThreadName = ""
	f.AssistantID = ""
	f.Thread = nil
}

func (f *Frame) Select(threadID, threadName string) {
	f.ThreadID = threadID
	f.ThreadName = threadName
	f.Thread = nil
}
