package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	// Sesuaikan path import sesuai dengan nama modul Anda
	httpHandler "github.com/RivaelManurung/go-webservices-clean-arch/internal/handler/http"
	"github.com/RivaelManurung/go-webservices-clean-arch/internal/repository/impl"
	"github.com/RivaelManurung/go-webservices-clean-arch/internal/usecase"
)

// Server adalah struct utama untuk aplikasi kita.
type Server struct {
	router *gin.Engine
}

// NewServer membuat instance server baru dengan semua dependensi dan rute yang sudah di-setup.
func NewServer() *Server {
	// Variabel Konfigurasi
	jwtSecret := "RAHASIA_NEGARA_JANGAN_DIBOCORKAN"

	// ================== DEPENDENCY INJECTION ==================
	// Inisialisasi semua komponen dari lapisan terluar ke terdalam.

	// --- User Dependencies ---
	userRepo := impl.NewInMemoryUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo, jwtSecret)
	userHandler := httpHandler.NewUserHandler(userUsecase)

	// // --- Product Dependencies ---
	// productRepo := impl.NewInMemoryProductRepository()
	// productUsecase := usecase.NewProductUsecase(productRepo)
	// productHandler := httpHandler.NewProductHandler(productUsecase)
	// // ==========================================================

	// Inisialisasi router Gin
	router := gin.Default()

	// ===================== REGISTER ROUTES ======================
	// Rute publik (tidak perlu login)
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	// Grup rute yang dilindungi (perlu token JWT)
	api := router.Group("/api")
	api.Use(httpHandler.AuthMiddleware(jwtSecret))
	{
		// api.POST("/products", productHandler.CreateProduct)
		// api.GET("/products/:id", productHandler.GetProductByID)
	}
	// ==========================================================

	server := &Server{
		router: router,
	}

	return server
}

// Start menjalankan server HTTP pada port tertentu.
func (s *Server) Start(addr string) error {
	fmt.Printf("🚀 Server berjalan di http://localhost%s\n", addr)
	return s.router.Run(addr)
}
