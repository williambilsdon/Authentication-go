package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/williambilsdon/authentication-go/internal/authuser"
)

func Initialise() *http.Server {
	r := handlers()

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return server
}

func handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login", authuser.LoginHandler)

	return r
}
