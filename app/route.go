package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Server struct of server
type Server struct {
	Env    string
	Router *mux.Router
}

// Routes do create route
func (s *Server) Routes() {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/api", s.handleSomething())
	s.Router.HandleFunc("/api/user", s.middlewareSomething(s.handleSomething()))
}

// Run do start server
func (s *Server) Run() {
	srv := &http.Server{
		Handler: s.Router,
		Addr:    ":80",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
