package main

import (
	"go-webservices-clean-arch/cmd/internal/server" // ✅ ini benar sekarang
)

func main() {
	app := server.NewServer()

	app.Run()
}
