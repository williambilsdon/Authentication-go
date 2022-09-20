package authapi

import (
	"database/sql"
	"log"
)

func (s *Server) newDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/users")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return db

}

func (s *Server) CreateUser(params *Register) error {
	_, err := s.db.Query("INSERT INTO users (Firstname, Lastname, Username, Password) VALUES (?, ?, ?, ?)", params.Firstname, params.Lastname, params.Username, params.Password)
	if err != nil {
		return err
	}

	return nil
}
