package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/williambilsdon/authentication-go/internal/authapi"
	repo "github.com/williambilsdon/authentication-go/internal/repository"
	"github.com/williambilsdon/authentication-go/internal/server"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/workout")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	authRepo := repo.NewAuthRepo(db)
	authService := authapi.NewAuthService(authRepo)
	authController := authapi.NewAuthController(authService)

	server := server.NewServer(authController)

	server.Start()

}
