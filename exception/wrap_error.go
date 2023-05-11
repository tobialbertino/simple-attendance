package exception

import "fmt"

// Custom wrap error
type WrappedError struct {
	Code    int
	Context string
	Err     error
}

func (w *WrappedError) Error() string {
	return fmt.Sprintf(`%s: %v, %v `, w.Context, w.Code, w.Err)
}

func Wrap(contextInfo string, code int, err error) *WrappedError {
	return &WrappedError{
		Context: contextInfo,
		Code:    code,
		Err:     err,
	}
}
