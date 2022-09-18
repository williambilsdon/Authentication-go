package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Login struct {
	username string
	password string
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	var loginBody Login
	err := json.NewDecoder(req.Body).Decode(&loginBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Print(loginBody.password)

	w.Write([]byte(loginBody.username))
}

func main() {

	http.HandleFunc("/login", loginHandler)

	fmt.Print("Serving")
	http.ListenAndServe(":8080", nil)
}
