package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kibreab/backend/internal/handlers"
)

func Setup(r *gin.Engine, 
	authHandler *handlers.AuthHandler, 
	deckHandler *handlers.DeckHandler, 
	flashHandler *handlers.FlashcardHandler,
	testHandler *handlers.TestHandler,
	) {
	api := r.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// Flashcard routes
	decks := api.Group("/decks")
	{
		decks.POST("/", deckHandler.CreateDeck)
		decks.GET("/", deckHandler.ListDecks)
		decks.DELETE("/:id", deckHandler.DeleteDeck)
		decks.PATCH("/:id", deckHandler.UpdateDeckTitle)

		decks.POST("/:id/flashcards", flashHandler.Generate)
		decks.GET("/:id/flashcards", flashHandler.List)
	}

	// Flashcard specific routes
	cards := api.Group("/flashcards")
	{
		cards.GET("/:id", flashHandler.Get)
		cards.DELETE("/:id", flashHandler.Delete)
	}

	api.GET("/test-openai", testHandler.CheckOpenAIKey)
}
