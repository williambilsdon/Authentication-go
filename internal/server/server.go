package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/williambilsdon/authentication-go/internal/authuser"
)

type Server struct {
	S *http.Server
}

func Initialise() *Server {
	r := mux.NewRouter()
	r.HandleFunc("/login", authuser.LoginHandler)
	server := Server{
		S: &http.Server{
			Addr:    ":8080",
			Handler: r,
		},
	}

	return &server
}
