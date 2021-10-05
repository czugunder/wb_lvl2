package calendar

import (
	"math/rand"
	"sync"
	"time"
)

// Event - тип, описывающий событие в календаре
type Event struct {
	Name    string    `json:"name"`
	Place   string    `json:"place"`
	UserUID string    `json:"user_uid"`
	StartDT time.Time `json:"start_dt"`
	EndDT   time.Time `json:"end_dt"`
}

// NewEvent создает экземпляр Event
func NewEvent() *Event {
	return &Event{}
}

// Calendar - тип, описывающий календарь
type Calendar struct {
	events map[string]*Event
	mu     *sync.RWMutex
}

// NewCalendar создает экземпляр Calendar
func NewCalendar() *Calendar {
	return &Calendar{
		events: make(map[string]*Event),
		mu:     &sync.RWMutex{},
	}
}

// CreateEvent создает событие в календаре
func (c *Calendar) CreateEvent(eve *Event) string {
	c.mu.Lock()
	eveUID := c.newUID()
	c.events[eveUID] = eve
	c.mu.Unlock()
	return eveUID
}

func (c *Calendar) newUID() string { // вызывается под Lock, поэтому здесь ничего не нужно
	var attempt string
	for {
		attempt = c.UIDContender()
		if _, found := c.events[attempt]; !found {
			break
		}
	}
	return attempt
}

// UIDContender создает вариант UID
func (c *Calendar) UIDContender() string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, 32)
	var t int
	for i := range bytes {
		t = rand.Intn(3)
		if t == 0 {
			bytes[i] = byte(rand.Intn(10) + 48) // цифра
		}
		if t == 1 {
			bytes[i] = byte(rand.Intn(26) + 97) // буква малая
		}
		if t == 2 {
			bytes[i] = byte(rand.Intn(26) + 65) // буква заглавная
		}
	}
	return string(bytes)
}

// UpdateEvent обновляет событие в календаре
func (c *Calendar) UpdateEvent(UID string, eve *Event) (*Event, error) {
	if storedEve, err := c.GetEvent(UID); err != nil {
		return nil, err
	} else {
		if eve.Name != "" {
			storedEve.Name = eve.Name
		}
		if eve.Place != "" {
			storedEve.Place = eve.Place
		}
		if eve.UserUID != "" {
			storedEve.UserUID = eve.UserUID
		}
		if !eve.StartDT.IsZero() {
			storedEve.StartDT = eve.StartDT
		}
		if !eve.EndDT.IsZero() {
			storedEve.EndDT = eve.EndDT
		}
		return storedEve, nil
	}
}

// DeleteEvent удаляет событие в календаре
func (c *Calendar) DeleteEvent(UID string) (eve *Event, err error) {
	c.mu.Lock()
	v, found := c.events[UID]
	if !found {
		err = &EventNotFound{}
	} else {
		eve = v
		delete(c.events, UID)
	}
	c.mu.Unlock()
	return
}

// EventsForDay возвращает список событий за указанный день для определенного пользователя
func (c *Calendar) EventsForDay(userUID string, day time.Time) []*Event {
	var res []*Event
	c.mu.RLock()
	for _, v := range c.events {
		if v.UserUID == userUID {
			if middle(v.StartDT.Year(), day.Year(), v.EndDT.Year()) &&
				middle(int(v.StartDT.Month()), int(day.Month()), int(v.EndDT.Month())) &&
				middle(v.StartDT.Day(), day.Day(), v.EndDT.Day()) {
				res = append(res, v)
			}
		}
	}
	c.mu.RUnlock()
	return res
}

// EventsForWeek возвращает список событий за указанную неделю для определенного пользователя
func (c *Calendar) EventsForWeek(userUID string, year time.Time) []*Event {
	var res []*Event
	c.mu.RLock()
	var y1, w1, yx, wx, y2, w2 int
	for _, v := range c.events {
		if v.UserUID == userUID {
			y1, w1 = v.StartDT.ISOWeek()
			yx, wx = year.ISOWeek()
			y2, w2 = v.EndDT.ISOWeek()
			if middle(y1, yx, y2) && middle(w1, wx, w2) {
				res = append(res, v)
			}
		}
	}
	c.mu.RUnlock()
	return res
}

// EventsForMonth возвращает список событий за указанный месяц для определенного пользовател
func (c *Calendar) EventsForMonth(userUID string, month time.Time) []*Event {
	var res []*Event
	c.mu.RLock()
	for _, v := range c.events {
		if v.UserUID == userUID {
			if middle(v.StartDT.Year(), month.Year(), v.EndDT.Year()) &&
				middle(int(v.StartDT.Month()), int(month.Month()), int(v.EndDT.Month())) {
				res = append(res, v)
			}
		}
	}
	c.mu.RUnlock()
	return res
}

func middle(first, x, second int) bool {
	if first <= x && x <= second {
		return true
	}
	return false
}

// GetEvent возвращает событие по его UID
func (c *Calendar) GetEvent(UID string) (eve *Event, err error) {
	c.mu.RLock()
	if v, found := c.events[UID]; !found {
		err = &EventNotFound{}
	} else {
		eve = v
	}
	c.mu.RUnlock()
	return
}
