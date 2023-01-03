package server

import (
	"context"
	"fmt"
	"io"
	"isaevfeed/notifier/internal/notifier"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Addr          string
	Router        *mux.Router
	NotifierEmail *notifier.NotifierEmail
}

func New() *Server {
	host, _ := os.LookupEnv("SERVER_HOST")
	port, _ := os.LookupEnv("SERVER_PORT")

	return &Server{Addr: fmt.Sprintf("%s:%s", host, port), Router: mux.NewRouter(), NotifierEmail: notifier.New(false)}
}

func (s *Server) Start() {
	srv := http.Server{
		Handler: s.Router,
		Addr:    s.Addr,
	}

	s.Router.HandleFunc("/hello", s.HandleHello()).Methods("GET")
	s.Router.HandleFunc("/api/v1/send-email", s.HandleSendEmail()).Methods("POST")

	go func() {
		srv.ListenAndServe()
		log.Print("Server is stopped")
	}()

	log.Print("Server is started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}

func (s *Server) HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Service's working normally")
	}
}

func (s *Server) HandleSendEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		email := req.FormValue("email")

		if email == "" {
			io.WriteString(w, MakeResponse(http.StatusBadRequest, "email - обязательное поле"))
			return
		}

		err := s.NotifierEmail.Send(email)
		if err != nil {
			io.WriteString(w, MakeResponse(http.StatusInternalServerError, fmt.Sprintf("Ошибка на сервере: %s", err)))
			return
		}

		io.WriteString(w, MakeResponse(http.StatusOK, "Задача на отправку успешно создана"))
	}
}
