package exception

import "fmt"

// Custom wrap error
type ClientError struct {
	Code    int
	Message string
}

func (w *ClientError) Error() string {
	return fmt.Sprintf(`%v: %v `, w.Code, w.Message)
}

func NewClientError(message string, code int) *ClientError {
	return &ClientError{
		Code:    code,
		Message: message,
	}
}
