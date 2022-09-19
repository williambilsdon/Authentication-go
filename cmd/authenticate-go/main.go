package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Login struct {
	Username string
	Password string
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	var loginBody Login
	err := json.NewDecoder(req.Body).Decode(&loginBody)
	if err != nil {
		fmt.Print(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Printf("Username: %s\nPassword: %s\n", loginBody.Username, loginBody.Password)
	w.Write([]byte(loginBody.Username))
}

func main() {

	http.HandleFunc("/login", loginHandler)

	fmt.Print("Serving")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
