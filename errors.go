package gerr

import (
	"fmt"
	"runtime"
)

func Wrap(error error) WrappedError {
	return WrapWith(error, 0, "", "")
}

func WrapWith(error error, errorCode int, userMsg string, internalMsg string) WrappedError {
	if e, ok := error.(*err); ok {
		newTrace := captureStackTrace()
		return &err{
			error:       e,
			stackTrace:  append(newTrace, e.stackTrace...),
			errorCode:   e.errorCode,
			userMessage: userMsg,
			internalMsg: internalMsg,
		}
	}

	return &err{
		error:       error,
		errorCode:   errorCode,
		userMessage: userMsg,
		internalMsg: internalMsg,
		stackTrace:  captureStackTrace(),
	}
}

func captureStackTrace() []string {
	var trace []string
	for i := 2; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			break
		}
		trace = append(trace, fmt.Sprintf("%s:%d %s", file, line, fn.Name()))
	}
	return trace
}
