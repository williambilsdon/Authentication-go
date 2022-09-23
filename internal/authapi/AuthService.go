package authapi

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/williambilsdon/authentication-go/internal/models"
	repo "github.com/williambilsdon/authentication-go/internal/repository"
)

type AuthService interface {
	CreateUser(user models.User) error
	Login(userLogin models.UserLogin) error
}

type authService struct {
	r repo.AuthRepo
}

func NewAuthService(repo repo.AuthRepo) *authService {
	return &authService{
		r: repo,
	}
}

func (s *authService) CreateUser(user models.User) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 4)
	if err != nil {
		return errors.New("failed to encrypt password")
	}

	err = s.r.CreateUser(user, encryptedPassword)

	return err
}

func (s *authService) Login(userLogin models.UserLogin) error {
	result := s.r.Login(userLogin.Username)
	var resultBytes []byte
	err := result.Scan(&resultBytes)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(resultBytes, []byte(userLogin.Password))
	if err != nil {
		return err
	}

	return nil
}
