package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/williambilsdon/authentication-go/internal/authapi"
)

func main() {

	server := authapi.Server{}
	server = server.NewServer()

	fmt.Print("Serving")
	server.Start()
}
