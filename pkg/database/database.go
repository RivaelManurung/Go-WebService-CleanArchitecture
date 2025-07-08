package database

import (
	"log"

	"go-webservices-clean-arch/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialize(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Running database migrations...")
	if err := db.AutoMigrate(&user.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db
}