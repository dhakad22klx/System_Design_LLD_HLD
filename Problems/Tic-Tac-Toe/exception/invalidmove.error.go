package exception

type InvalidMoveError struct {
	Message string
}

func NewInvalidMoveError(message string) InvalidMoveError {
	return InvalidMoveError{Message: message}
}

func (exception InvalidMoveError) Error() string {
	return exception.Message
}
