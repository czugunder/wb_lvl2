package server

// IncorrectMethod - ошибка, неверный метод
type IncorrectMethod struct{}

func (e *IncorrectMethod) Error() string {
	return "incorrect method was used for this request"
}

// InvalidInput - ошибка, неправильные входные данные
type InvalidInput struct{}

func (e *InvalidInput) Error() string {
	return "invalid input"
}

// InvalidEventUID - ошибка, нет события с таким UID
type InvalidEventUID struct{}

func (e *InvalidEventUID) Error() string {
	return "invalid event UID"
}

// InvalidDate - ошибка, неверная дата
type InvalidDate struct{}

func (e *InvalidDate) Error() string {
	return "invalid date"
}
