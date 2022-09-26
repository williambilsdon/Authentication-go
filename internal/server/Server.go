package server

import (
	"log"
	"net/http"

	"github.com/williambilsdon/authentication-go/internal/authapi"
)

type Server struct {
	auth       authapi.AuthController
	httpServer *http.Server
}

func NewServer(auth authapi.AuthController) Server {
	server := Server{
		auth: auth,
	}

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: server.createHandler(),
	}

	server.httpServer = &httpServer

	return server
}

func (s *Server) createHandler() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", s.auth.Login)
	mux.HandleFunc("/register", s.auth.CreateUser)
	mux.HandleFunc("/dosomething", s.auth.DoSomething)

	return mux
}

func (s *Server) Start() {
	log.Fatal(s.httpServer.ListenAndServe())
}
