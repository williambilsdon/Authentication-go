package authuser

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type login struct {
	Username string
	Password string
}

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	var loginBody login
	err := json.NewDecoder(req.Body).Decode(&loginBody)
	if err != nil {
		fmt.Print(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Printf("Username: %s\nPassword: %s\n", loginBody.Username, loginBody.Password)
	w.Write([]byte(loginBody.Username))
}
