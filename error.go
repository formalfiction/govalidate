package validate

type Error struct {
	message string
}

func NewError(m string) Error {
	return Error{m}
}

func (e Error) Error() string {
	return e.message
}
