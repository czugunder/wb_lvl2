package calendar

type EventNotFound struct{}

func (e *EventNotFound) Error() string {
	return "event with such UID was not found"
}
