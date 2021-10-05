package server

import (
	"encoding/json"
	"net/http"
	"time"
	"wb_lvl2/develop/dev11/internal/dev11/calendar"
)

var (
	errMethod   = &IncorrectMethod{}
	errDate     = &InvalidDate{}
	errInput    = &InvalidInput{}
	errEventUID = &InvalidEventUID{}
)

type eventUpdateRAW struct {
	UID     string `json:"uid"`
	Name    string `json:"name"`
	Place   string `json:"place"`
	UserUID string `json:"user_uid"`
	StartDT string `json:"start_dt"`
	EndDT   string `json:"end_dt"`
}

func newEventUpdateRAW() *eventUpdateRAW {
	return &eventUpdateRAW{}
}

func (e *eventUpdateRAW) convertToEvent() *calendar.Event {
	eve := calendar.NewEvent()
	eve.Name = e.Name
	eve.Place = e.Place
	eve.UserUID = e.UserUID
	eve.StartDT, _ = time.Parse(time.RFC3339, e.StartDT)
	eve.EndDT, _ = time.Parse(time.RFC3339, e.EndDT)
	return eve
}

func (e *eventUpdateRAW) valid() bool {
	if e.UID == "" {
		return false
	}
	return true
}

type eventAddRAW struct {
	Name    string `json:"name"`
	Place   string `json:"place"`
	UserUID string `json:"user_uid"`
	StartDT string `json:"start_dt"`
	EndDT   string `json:"end_dt"`
}

func newEventAddRAW() *eventAddRAW {
	return &eventAddRAW{}
}

func (e *eventAddRAW) convertToEvent() *calendar.Event {
	eve := calendar.NewEvent()
	eve.Name = e.Name
	eve.Place = e.Place
	eve.UserUID = e.UserUID
	eve.StartDT, _ = time.Parse(time.RFC3339, e.StartDT)
	eve.EndDT, _ = time.Parse(time.RFC3339, e.EndDT)
	return eve
}

func (e *eventAddRAW) valid() bool {
	if e.UserUID == "" {
		return false
	}
	if _, err := time.Parse(time.RFC3339, e.StartDT); err != nil {
		return false
	}
	if _, err := time.Parse(time.RFC3339, e.EndDT); err != nil {
		return false
	}
	return true
}

func (s *Server) addEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.response(false, w, errMethod.Error(), http.StatusServiceUnavailable)
		return
	}
	eveRaw := newEventAddRAW()
	err := json.NewDecoder(r.Body).Decode(&eveRaw)
	if err != nil {
		s.response(false, w, errInput.Error(), http.StatusBadRequest)
		return
	}
	if !eveRaw.valid() {
		s.response(false, w, errInput.Error(), http.StatusBadRequest)
		return
	}
	eve := eveRaw.convertToEvent()
	uid := s.calendar.CreateEvent(eve)
	s.response(true, w, struct {
		UID string `json:"uid"`
	}{UID: uid}, http.StatusOK)
}

func (s *Server) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.response(false, w, errMethod.Error(), http.StatusServiceUnavailable)
		return
	}
	eveRaw := newEventUpdateRAW()
	err := json.NewDecoder(r.Body).Decode(&eveRaw)
	if err != nil {
		s.response(false, w, errInput.Error(), http.StatusBadRequest)
		return
	}
	if !eveRaw.valid() {
		s.response(false, w, errInput.Error(), http.StatusBadRequest)
		return
	}
	eve := eveRaw.convertToEvent()
	curEve, errU := s.calendar.UpdateEvent(eveRaw.UID, eve)
	if errU != nil {
		s.response(false, w, nil, http.StatusBadRequest)
	} else {
		s.response(true, w, curEve, http.StatusOK)
	}
}

func (s *Server) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.response(false, w, errMethod.Error(), http.StatusServiceUnavailable)
		return
	}
	eveRaw := newEventUpdateRAW()
	err := json.NewDecoder(r.Body).Decode(&eveRaw)
	if err != nil {
		s.response(false, w, errInput.Error(), http.StatusBadRequest)
		return
	}
	if !eveRaw.valid() {
		s.response(false, w, errInput.Error(), http.StatusBadRequest)
		return
	}
	delEve, errU := s.calendar.DeleteEvent(eveRaw.UID)
	if errU != nil {
		s.response(false, w, errEventUID.Error(), http.StatusBadRequest)
	} else {
		s.response(true, w, delEve, http.StatusOK)
	}
}

func (s *Server) dayEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.response(false, w, errMethod.Error(), http.StatusServiceUnavailable)
		return
	}
	user, foundUser := r.URL.Query()["user_id"]
	dateRaw, foundDate := r.URL.Query()["date"]
	if !foundUser || !foundDate {
		s.response(false, w, errInput.Error(), http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02", dateRaw[0])
	if err != nil {
		s.response(false, w, errDate.Error(), http.StatusBadRequest)
		return
	}
	eves := s.calendar.EventsForDay(user[0], date)
	s.response(true, w, eves, http.StatusOK)
}

func (s *Server) weekEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.response(false, w, errMethod.Error(), http.StatusServiceUnavailable)
		return
	}
	user, foundUser := r.URL.Query()["user_id"]
	dateRaw, foundDate := r.URL.Query()["date"]
	if !foundUser || !foundDate {
		s.response(false, w, errInput.Error(), http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02", dateRaw[0])
	if err != nil {
		s.response(false, w, errDate.Error(), http.StatusBadRequest)
		return
	}
	eves := s.calendar.EventsForWeek(user[0], date)
	s.response(true, w, eves, http.StatusOK)
}

func (s *Server) monthEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.response(false, w, errMethod.Error(), http.StatusServiceUnavailable)
		return
	}
	user, foundUser := r.URL.Query()["user_id"]
	dateRaw, foundDate := r.URL.Query()["date"]
	if !foundUser || !foundDate {
		s.response(false, w, errInput.Error(), http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02", dateRaw[0])
	if err != nil {
		s.response(false, w, errDate.Error(), http.StatusBadRequest)
		return
	}
	eves := s.calendar.EventsForMonth(user[0], date)
	s.response(true, w, eves, http.StatusOK)
}

// Result тип для ответа, имеет поле result
type Result struct {
	Result interface{} `json:"result"`
}

// Error тип для ответа, имеет поле error
type Error struct {
	Error interface{} `json:"error"`
}

func (s *Server) response(isResult bool, w http.ResponseWriter, payload interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	if isResult {
		payload = Result{Result: payload}
	} else {
		payload = Error{Error: payload}
	}
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "response error", http.StatusInternalServerError)
	}

}
