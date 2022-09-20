package authapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Register struct {
	Firstname string
	Lastname  string
	Username  string
	Password  string
}

func (s *Server) RegisterHandler(w http.ResponseWriter, req *http.Request) {
	var registerBody Register
	err := json.NewDecoder(req.Body).Decode(&registerBody)
	if err != nil {
		fmt.Print(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Printf("Username: %s\nPassword: %s\n", registerBody.Username, registerBody.Password)
	w.Write([]byte(registerBody.Username))
}
