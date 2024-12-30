# Gerr: Go Error Wrapper Library

**Gerr** is a library for wrapping errors with detailed information, including:
- Custom error codes.
- User-friendly error messages.
- Internal messages for debugging.
- A full stack trace to track the origin of the error.

This library is designed to make error handling more informative and structured, improving both debugging and user communication.

## Features

- **Error wrapping:** Chain errors to preserve context.
- **Stack trace:** Capture stack traces at each wrapping point.
- **Custom error codes:** Add meaningful error codes.
- **User-friendly messages:** Provide messages intended for end-users.
- **Internal messages:** Include detailed technical descriptions for debugging.
- **Root error identification:** Access the original error in a chain.

---

## Installation

```bash
go get github.com/akselarzuman/gerr
```

---

## Usage

### Example: Wrapping and Accessing Errors

```go
package main

import (
	"errors"
	"fmt"

	"github.com/akselarzuman/gerr"
)

func main() {
	// Create an initial error
	baseErr := gerr.WrapWith(errors.New("database connection failed"), 503, "Service temporarily unavailable.", "Unable to connect to DB at db.example.com")

	// Wrap the error with additional context
	wrappedErr := gerr.WrapWith(baseErr, 503, baseErr.UserMessage(), "Failed during user login attempt")

	// Print the root error with its full details
	fmt.Println("Root Error:")
	fmt.Println(wrappedErr.RootError().FullError())

	// Print the fully wrapped error with its full stack trace
	fmt.Println("\nWrapped Error:")
	fmt.Println(wrappedErr.FullError())
}
```

### Output Example:

```plaintext
Root Error:
Error: database connection failed
Error Type: 503
User Message: Service temporarily unavailable.
Internal Message: Unable to connect to DB at db.example.com
Stack Trace:
	/path/to/file.go:10 main.main

Wrapped Error:
Error: Failed during user login attempt
Error Type: 503
User Message: Service temporarily unavailable.
Internal Message: Failed during user login attempt
Stack Trace:
	/path/to/file.go:15 main.main
	/path/to/file.go:10 main.main
```

---

## API Documentation

### Wrapping Errors

#### `gerr.Wrap(err error) WrappedError`
Wraps an existing error, capturing the stack trace.

#### `gerr.WrapWith(err error, errorCode int, userMsg string, internalMsg string) WrappedError`
Wraps an error with additional information:
- `errorCode`: A custom error code.
- `userMsg`: A message intended for the end user.
- `internalMsg`: A message intended for developers.

### Accessing Error Details

#### `Error() string`
Returns the error message.

#### `ErrorCode() int`
Returns the custom error code.

#### `UserMessage() string`
Returns the user-friendly error message.

#### `InternalMessage() string`
Returns the internal debugging message.

#### `StackTrace() []string`
Returns the captured stack trace.

#### `FullError() string`
Returns a detailed error report including the error message, code, user message, internal message, and stack trace.

#### `RootError() WrappedError`
Returns the root cause of the error.

---

## License

This library is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

---

## Contributions

Contributions are welcome! Feel free to open issues or submit pull requests.