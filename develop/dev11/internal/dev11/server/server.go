package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wb_lvl2/develop/dev11/internal/dev11/calendar"
)

type server struct {
	calendar *calendar.Calendar
	config   *config
}

func NewSever(config *config) *server {
	return &server{
		config:   config,
		calendar: calendar.NewCalendar(),
	}
}

func logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeBefore := time.Now()
		handler.ServeHTTP(w, r)
		log.Printf("%s\t%s\t%s\t%s\n", r.Method, r.URL, timeBefore, time.Since(timeBefore))
	})
}

func (s *server) handle() {
	http.Handle("/create_event", logger(http.HandlerFunc(s.addEvent)))
	http.Handle("/update_event", logger(http.HandlerFunc(s.updateEvent)))
	http.Handle("/delete_event", logger(http.HandlerFunc(s.deleteEvent)))
	http.Handle("/events_for_day", logger(http.HandlerFunc(s.dayEvents)))
	http.Handle("/events_for_week", logger(http.HandlerFunc(s.weekEvents)))
	http.Handle("/events_for_month", logger(http.HandlerFunc(s.monthEvents)))
}

func (s *server) runServer(err chan error) {
	go func() {
		err <- http.ListenAndServe(s.config.address, nil) // дефолтный http.DefaultServeMux
	}()
}

func (s *server) Run() {
	s.handle()
	sigint := make(chan os.Signal)
	errors := make(chan error)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM) // ctrl+c или kill <pid>
	s.runServer(errors)
	select {
	case <-sigint:
		log.Println("server stopped")
		return
	case err := <-errors:
		log.Println(err)
	}
}
