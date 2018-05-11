package app

import "net/http"

func (s *Server) handleSomething() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
		w.WriteHeader(200)
	}
}
