package authapi

import (
	"database/sql"
	"log"
	"net/http"
)

type Server struct {
	db *sql.DB
	s  *http.Server
}

func (s *Server) NewServer() Server {
	db := s.newDatabase()

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: s.createHandler(),
	}

	return Server{
		db: db,
		s:  &httpServer,
	}

}

func (s *Server) createHandler() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", s.LoginHandler)
	mux.HandleFunc("/register", s.RegisterHandler)

	return mux
}

func (s *Server) Start() {
	log.Fatal(s.s.ListenAndServe())
}
