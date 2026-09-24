package gerr_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/akselarzuman/gerr"
)

func TestHasCode(t *testing.T) {
	const customCode gerr.ErrorCode = 1001
	notFound := gerr.WrapWith(errors.New("missing record"), int(gerr.NotFoundError), "Not found", "record missing")
	outer := gerr.WrapWith(fmt.Errorf("repository: %w", notFound), int(gerr.InternalError), "Request failed", "service failed")

	tests := []struct {
		name string
		err  error
		code gerr.ErrorCode
		want bool
	}{
		{name: "nil", code: gerr.UnknownError},
		{name: "plain error", err: errors.New("plain"), code: gerr.UnknownError},
		{name: "matching code", err: notFound, code: gerr.NotFoundError, want: true},
		{name: "different code", err: notFound, code: gerr.ValidationError},
		{name: "standard wrapper", err: fmt.Errorf("service: %w", notFound), code: gerr.NotFoundError, want: true},
		{name: "gerr wrapper", err: gerr.Wrap(notFound), code: gerr.NotFoundError, want: true},
		{name: "custom code", err: gerr.WrapWith(errors.New("custom"), int(customCode), "", ""), code: customCode, want: true},
		{name: "outer code takes precedence", err: outer, code: gerr.InternalError, want: true},
		{name: "inner code does not match", err: outer, code: gerr.NotFoundError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gerr.HasCode(tt.err, tt.code); got != tt.want {
				t.Errorf("HasCode(%v, %d) = %v, want %v", tt.err, tt.code, got, tt.want)
			}
		})
	}
}
