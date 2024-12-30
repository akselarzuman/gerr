package gerr

type ErrorCode int

const (
	UnknownError ErrorCode = iota
	ValidationError
	InternalError
	NotFoundError
	UnauthorizedError
	ForbiddenError
)
