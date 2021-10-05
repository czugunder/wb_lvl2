package cut

type rangeError struct{}

func (e *rangeError) Error() string {
	return "only S, S-, -F, S-F range formats allowed, count starts with 1"
}

func NewRangeError() *rangeError {
	return &rangeError{}
}
