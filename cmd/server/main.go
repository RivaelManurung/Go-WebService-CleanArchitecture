package main

import (
	"go-webservices-clean-arch/cmd/internal/server" 
)

func main() {
	app := server.NewServer()

	app.Run()
}
