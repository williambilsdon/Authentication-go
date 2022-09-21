package authapi

import (
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
	err := s.r.CreateUser(user)

	return err
}

func (s *authService) Login(userLogin models.UserLogin) error {
	result := s.r.Login(userLogin)
	var resultBytes []byte
	err := result.Scan(&resultBytes)
	if err != nil {
		return err
	}

	return nil
}
