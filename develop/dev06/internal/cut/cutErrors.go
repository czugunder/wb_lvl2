package cut

// RangeError ошибка флага -f
type RangeError struct{}

func (e *RangeError) Error() string {
	return "only S, S-, -F, S-F range formats allowed, count starts with 1"
}

// NewRangeError создает экземпляр RangeError
func NewRangeError() *RangeError {
	return &RangeError{}
}
