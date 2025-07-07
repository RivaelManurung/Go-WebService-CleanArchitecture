package main

import "log"

func main() {
	// 1. Buat server baru yang sudah dikonfigurasi dari server.go
	server := NewServer()

	// 2. Jalankan server
	err := server.Start(":8080")
	if err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}