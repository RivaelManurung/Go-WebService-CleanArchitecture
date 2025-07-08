package server

import (
	"log"

	"go-webservices-clean-arch/config"
	"go-webservices-clean-arch/pkg/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Server adalah struct utama aplikasi kita
type Server struct {
	config *config.Config
	db     *gorm.DB
	router *gin.Engine
}

// NewServer membuat instance server baru setelah inisialisasi
func NewServer() *Server {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db := database.Initialize(cfg.DatabaseURL)

	server := &Server{
		config: cfg,
		db:     db,
		router: gin.Default(),
	}

	// Daftarkan semua rute
	server.setupRoutes()

	return server
}

// Run menjalankan server HTTP
func (s *Server) Run() {
	log.Println("Starting server on port", s.config.ServerPort)
	if err := s.router.Run(s.config.ServerPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}