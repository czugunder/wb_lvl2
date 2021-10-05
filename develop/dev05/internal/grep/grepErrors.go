package grep

type NoPatternError struct{}

func (e *NoPatternError) Error() string {
	return "no pattern was provided in args"
}

func NewNoPatternError() *NoPatternError {
	return &NoPatternError{}
}
