package server

import (
	v1 "go-webservices-clean-arch/api/v1"
	"go-webservices-clean-arch/internal/user"
	"go-webservices-clean-arch/pkg/auth"
)

// setupRoutes menginisialisasi semua dependensi dan mendaftarkan rute
func (s *Server) setupRoutes() {
	// === Inisialisasi Dependensi ===
	userRepo := user.NewUserRepository(s.db)
	authService := auth.NewAuthService(s.config.JWTSecretKey)
	userService := user.NewUserService(userRepo, authService)
	userHandler := v1.NewUserHandler(userService)

	// === Pengelompokan Rute ===
	apiV1 := s.router.Group("/v1")

	// --- Rute Publik (Tidak perlu login) ---
	apiV1.POST("/register", userHandler.Register)
	apiV1.POST("/login", userHandler.Login)

	// --- Rute Terproteksi (Wajib login) ---
	authRoutes := apiV1.Group("/")
	authRoutes.Use(auth.Middleware(authService)) // Terapkan middleware di sini
	{
		authRoutes.GET("/me", userHandler.GetProfile)
		// Tambahkan rute terproteksi lainnya di sini
		// Misalnya: authRoutes.POST("/posts", postHandler.Create)
	}
}
