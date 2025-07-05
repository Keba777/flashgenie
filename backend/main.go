package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/kibreab/backend/config"
	"github.com/kibreab/backend/internal/handlers"
	"github.com/kibreab/backend/internal/repository"
	"github.com/kibreab/backend/internal/services"
	"github.com/kibreab/backend/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	// Load configuration
	cfg := config.Load()

	// Build PostgreSQL DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	repo := repository.NewPostgresRepo(db)
	if err := repo.AutoMigrate(); err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	}

	// Initialize services
	authService := services.NewAuthService(repo, cfg.JWTSecret)
	aiService := services.NewAIService(cfg.TogetherAIKey)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	flashHandler := handlers.NewFlashcardHandler(aiService, repo)
	deckHandler := handlers.NewDeckHandler(repo)
	testHandler := handlers.NewTestHandler(aiService)

	// Setup Gin router and routes
	r := gin.Default()
	routes.Setup(r, authHandler, deckHandler, flashHandler, testHandler, cfg.JWTSecret)

	// Start server on the port Render provides (or fallback)
	log.Printf("🚀 Server running on port %d...\n", cfg.ServerPort)
	if err := r.Run(fmt.Sprintf(":%d", cfg.ServerPort)); err != nil {
		log.Fatalf("❌ Failed to run server: %v", err)
	}
}
