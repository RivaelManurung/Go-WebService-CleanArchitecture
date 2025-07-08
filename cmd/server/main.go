package main

import (
	"log"
	"github.com/RivaelManurung/go-webservices-clean-arch/api/v1"
	"github.com/RivaelManurung/go-webservices-clean-arch/config"
	"github.com/RivaelManurung/go-webservices-clean-arch/internal/user"
	"github.com/RivaelManurung/go-webservices-clean-arch/pkg/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 1. Muat Konfigurasi
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	if cfg.DatabaseURL == "" || cfg.JWTSecretKey == "" {
		log.Fatal("DATABASE_URL or JWT_SECRET_KEY is not set")
	}

	// 2. Inisialisasi Database
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 3. Jalankan Migrasi
	log.Println("Running database migrations...")
	if err := db.AutoMigrate(&user.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// 4. Inisialisasi Dependensi
	userRepo := user.NewUserRepository(db)
	authService := auth.NewAuthService(cfg.JWTSecretKey)
	userService := user.NewUserService(userRepo, authService)
	userHandler := v1.NewUserHandler(userService)

	// 5. Inisialisasi Router & Atur Rute
	router := gin.Default()
	apiV1 := router.Group("/v1")
	{
		apiV1.POST("/register", userHandler.Register)
		apiV1.POST("/login", userHandler.Login)
	}

	// 6. Jalankan Server
	log.Println("Starting server on port :8080")
	router.Run(":8080")
}