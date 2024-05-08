package exception

type PasswordIsWrongError struct {
	Error string
}

func NewPasswordIsWrongError(error string) PasswordIsWrongError {
	return PasswordIsWrongError{Error: error}
}
