package exceptions

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

type UnauthorizeError struct {
	Error string
}

func NewUnauthorizeError(error string) UnauthorizeError {
	return UnauthorizeError{Error: error}
}
