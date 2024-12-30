package gerr

import (
	"fmt"
	"strings"
)

var _ WrappedError = &err{}

type WrappedError interface {
	error
	ErrorCode() int
	UserMessage() string
	InternalMessage() string
	StackTrace() []string
	FullError() string
	RootError() WrappedError
}

type err struct {
	error       error
	stackTrace  []string
	errorCode   int
	userMessage string
	internalMsg string
}

func (e *err) Error() string {
	return fmt.Sprintf("%v", e.error)
}

func (e *err) FullError() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Error: %v\n", e.Error()))
	sb.WriteString(fmt.Sprintf("Error Type: %v\n", e.ErrorCode()))
	sb.WriteString(fmt.Sprintf("User Message: %s\n", e.UserMessage()))
	sb.WriteString(fmt.Sprintf("Internal Message: %s\n", e.InternalMessage()))
	sb.WriteString("Stack Trace:\n")
	for _, frame := range e.StackTrace() {
		sb.WriteString(fmt.Sprintf("\t%s\n", frame))
	}
	return sb.String()
}

func (e *err) ErrorCode() int {
	return e.errorCode
}

func (e *err) UserMessage() string {
	return e.userMessage
}

func (e *err) InternalMessage() string {
	return e.internalMsg
}

func (e *err) StackTrace() []string {
	return e.stackTrace
}

func (e *err) RootError() WrappedError {
	if e.error == nil {
		return e
	}

	if wrappedErr, ok := e.error.(WrappedError); ok {
		return wrappedErr.RootError()
	}

	return e
}
