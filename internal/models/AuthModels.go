package models

type User struct {
	Username  string
	Firstname string
	Lastname  string
	Password  string
}

type UserLogin struct {
	Username string
	Password string
}
