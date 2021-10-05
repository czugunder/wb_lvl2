package grep

// NoPatternError ошибка отсутствия паттерна в аргументах
type NoPatternError struct{}

func (e *NoPatternError) Error() string {
	return "no pattern was provided in args"
}

// NewNoPatternError создает экземпляр NoPatternError
func NewNoPatternError() *NoPatternError {
	return &NoPatternError{}
}
