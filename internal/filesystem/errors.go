package filesystem

type IsFileError struct {
	Message string
}

func (e *IsFileError) Error() string {
	return e.Message
}

var ErrIsFileError = &IsFileError{Message: "this path is a directory"}
