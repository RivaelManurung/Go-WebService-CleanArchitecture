package server

import (
	v1 "go-webservices-clean-arch/api/v1"
	"go-webservices-clean-arch/internal/user"
	"go-webservices-clean-arch/pkg/auth"
)

// setupRoutes menginisialisasi semua dependensi dan mendaftarkan rute
func (s *Server) setupRoutes() {
	// === Inisialisasi Dependensi ===
	// Disini adalah tempat "Dependency Injection" terjadi
	userRepo := user.NewUserRepository(s.db)
	authService := auth.NewAuthService(s.config.JWTSecretKey)
	userService := user.NewUserService(userRepo, authService)
	userHandler := v1.NewUserHandler(userService)

	// === Pengelompokan Rute ===
	apiV1 := s.router.Group("/v1")
	{
		// Rute publik
		apiV1.POST("/register", userHandler.Register)
		apiV1.POST("/login", userHandler.Login)

		// Rute yang memerlukan otentikasi bisa ditambahkan di grup lain
		// authRoutes := apiV1.Group("/")
		// authRoutes.Use(yourAuthMiddleware)
		// {
		// 	authRoutes.GET("/profile", userHandler.GetProfile)
		// }
	}
}