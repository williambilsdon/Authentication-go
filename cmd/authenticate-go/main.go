package main

import (
	"fmt"
	"log"

	"github.com/williambilsdon/authentication-go/internal/server"
)

func main() {

	server := server.Initialise()

	fmt.Print("Serving")
	log.Fatal(server.S.ListenAndServe())
}
