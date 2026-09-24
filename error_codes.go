package gerr

import "errors"

type ErrorCode int

const (
	UnknownError ErrorCode = iota
	ValidationError
	InternalError
	NotFoundError
	UnauthorizedError
	ForbiddenError
)

// HasCode reports whether the first WrappedError found by errors.As has code.
// It returns false when err is nil or does not contain a WrappedError.
func HasCode(err error, code ErrorCode) bool {
	var wrapped WrappedError
	return errors.As(err, &wrapped) && wrapped.ErrorCode() == int(code)
}
