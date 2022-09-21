package repo

import (
	"database/sql"

	"github.com/williambilsdon/authentication-go/internal/models"
)

type AuthRepo interface {
	CreateUser(params models.User) error
	Login(params models.UserLogin) *sql.Row
}

type authRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *authRepo {
	return &authRepo{db}
}

func (r *authRepo) CreateUser(params models.User) error {
	_, err := r.db.Query("INSERT INTO users (Username, Firstname, Lastname, Password) VALUES (?, ?, ?, ?)", params.Username, params.Firstname, params.Lastname, params.Password)
	return err
}

func (r *authRepo) Login(params models.UserLogin) *sql.Row {
	result := r.db.QueryRow("SELECT Username FROM users WHERE Username = ? AND Password = ?", params.Username, params.Password)
	return result
}
