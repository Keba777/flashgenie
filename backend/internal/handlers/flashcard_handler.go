package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kibreab/backend/internal/repository"
	"github.com/kibreab/backend/internal/services"
)

type FlashcardHandler struct {
	aiSvc *services.AIService
	repo  repository.Repository
}

func NewFlashcardHandler(aiSvc *services.AIService, repo repository.Repository) *FlashcardHandler {
	return &FlashcardHandler{aiSvc: aiSvc, repo: repo}
}

// Generate flashcards using AI and store them
func (h *FlashcardHandler) Generate(c *gin.Context) {
    deckID, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid deck ID"})
        return
    }

    var req struct{ Content string }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    log.Printf("Generating flashcards for deck %s with input text length %d", deckID.String(), len(req.Content))

    cards, err := h.aiSvc.GenerateFlashcards(req.Content)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate flashcards: " + err.Error()})
        return
    }

    // Save generated cards
    for i := range cards {
        cards[i].DeckID = deckID
        if err := h.repo.CreateFlashcard(&cards[i]); err != nil {
            log.Printf("Failed to save flashcard: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save flashcard: " + err.Error()})
            return
        }
    }

    c.JSON(http.StatusOK, cards)
}


// List flashcards in a deck
func (h *FlashcardHandler) List(c *gin.Context) {
	deckID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid deck ID"})
		return
	}

	cards, err := h.repo.GetFlashcardsByDeck(deckID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cards)
}

// Optional: Get a single flashcard by ID
func (h *FlashcardHandler) Get(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid flashcard ID"})
		return
	}

	card, err := h.repo.GetFlashcardByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "flashcard not found"})
		return
	}

	c.JSON(http.StatusOK, card)
}

// Optional: Delete a flashcard
func (h *FlashcardHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid flashcard ID"})
		return
	}

	if err := h.repo.DeleteFlashcard(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
