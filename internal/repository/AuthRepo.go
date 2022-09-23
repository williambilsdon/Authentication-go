package repo

import (
	"database/sql"

	"github.com/williambilsdon/authentication-go/internal/models"
)

type AuthRepo interface {
	CreateUser(params models.User, password []byte) error
	Login(username string) *sql.Row
}

type authRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *authRepo {
	return &authRepo{db}
}

func (r *authRepo) CreateUser(params models.User, password []byte) error {
	_, err := r.db.Query("INSERT INTO users (Username, Firstname, Lastname, Password) VALUES (?, ?, ?, ?)", params.Username, params.Firstname, params.Lastname, password)
	return err
}

func (r *authRepo) Login(username string) *sql.Row {
	result := r.db.QueryRow("SELECT Password FROM users WHERE Username = ?", username)
	return result
}
