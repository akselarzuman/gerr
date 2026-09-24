package gerr_test

import (
	"errors"
	"fmt"
	"io/fs"
	"testing"

	"github.com/akselarzuman/gerr"
)

func TestStandardErrorChain(t *testing.T) {
	cause := &fs.PathError{Op: "open", Path: "missing.txt", Err: fs.ErrNotExist}
	wrapped := gerr.WrapWith(cause, int(gerr.NotFoundError), "File not found", "open failed")
	chain := gerr.Wrap(fmt.Errorf("service: %w", wrapped))

	if got := errors.Unwrap(wrapped); got != cause {
		t.Errorf("Unwrap() = %v, want original cause %v", got, cause)
	}
	if !errors.Is(chain, fs.ErrNotExist) {
		t.Error("errors.Is did not find the original sentinel through mixed wrappers")
	}
	var pathError *fs.PathError
	if !errors.As(chain, &pathError) || pathError != cause {
		t.Error("errors.As did not recover the original typed error through mixed wrappers")
	}
}
