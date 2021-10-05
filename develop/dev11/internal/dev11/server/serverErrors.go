package server

type IncorrectMethod struct{}

func (e *IncorrectMethod) Error() string {
	return "incorrect method was used for this request"
}

type InvalidInput struct{}

func (e *InvalidInput) Error() string {
	return "invalid input"
}

type InvalidEventUID struct{}

func (e *InvalidEventUID) Error() string {
	return "invalid event UID"
}

type InvalidDate struct{}

func (e *InvalidDate) Error() string {
	return "invalid date"
}
