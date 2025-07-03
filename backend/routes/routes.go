package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kibreab/backend/internal/handlers"
	"github.com/kibreab/backend/internal/middleware"
)

func Setup(
	r *gin.Engine,
	authHandler *handlers.AuthHandler,
	deckHandler *handlers.DeckHandler,
	flashHandler *handlers.FlashcardHandler,
	testHandler *handlers.TestHandler,
	jwtSecret string,
) {
	api := r.Group("/api")

	// Public (no auth)
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}
	api.GET("/test-openai", testHandler.CheckOpenAIKey)

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.JWTMiddleware(jwtSecret))

	// Deck endpoints
	decks := protected.Group("decks")
	{
		decks.POST("/", deckHandler.CreateDeck)
		decks.GET("/", deckHandler.ListDecks)
		decks.DELETE("/:id", deckHandler.DeleteDeck)
		decks.PATCH("/:id", deckHandler.UpdateDeckTitle)

		decks.POST("/:id/flashcards", flashHandler.Generate)
		decks.GET("/:id/flashcards", flashHandler.List)
	}

	// Flashcard-specific endpoints
	cards := protected.Group("flashcards")
	{
		cards.GET("/:id", flashHandler.Get)
		cards.DELETE("/:id", flashHandler.Delete)
	}
}
