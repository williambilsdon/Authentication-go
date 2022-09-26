package authapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/williambilsdon/authentication-go/internal/models"
)

type AuthController interface {
	CreateUser(w http.ResponseWriter, req *http.Request)
	Login(w http.ResponseWriter, req *http.Request)
}

type authController struct {
	s AuthService
}

func NewAuthController(service AuthService) *authController {
	return &authController{
		s: service,
	}
}

func (c *authController) CreateUser(w http.ResponseWriter, req *http.Request) {
	var newUser models.User
	err := json.NewDecoder(req.Body).Decode(&newUser)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = c.s.CreateUser(newUser)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "Failed to register new user", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("New user succesfully created."))
}

func (c *authController) Login(w http.ResponseWriter, req *http.Request) {
	var userLogin models.UserLogin
	err := json.NewDecoder(req.Body).Decode(&userLogin)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	token, err := c.s.Login(userLogin)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "Login Failed", http.StatusBadRequest)
		return
	}
	w.Write([]byte(token))
	w.Write([]byte("Login successful."))
}
