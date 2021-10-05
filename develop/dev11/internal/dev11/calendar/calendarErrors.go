package calendar

// EventNotFound - ошибка, событие не найденно
type EventNotFound struct{}

func (e *EventNotFound) Error() string {
	return "event with such UID was not found"
}
