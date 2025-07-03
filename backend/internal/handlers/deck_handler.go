// internal/handlers/deck_handler.go
package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kibreab/backend/internal/models"
	"github.com/kibreab/backend/internal/repository"
)

type DeckHandler struct {
	repo repository.Repository
}

func NewDeckHandler(repo repository.Repository) *DeckHandler {
	return &DeckHandler{repo: repo}
}

// CreateDeck creates a new deck for the authenticated user
func (h *DeckHandler) CreateDeck(c *gin.Context) {
	var req struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve userID from context (set by JWT middleware)
	uidAny, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}
	uidStr, ok := uidAny.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID in context"})
		return
	}
	userID, err := uuid.Parse(uidStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID format"})
		return
	}

	deck := &models.Deck{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
	}
	if err := h.repo.CreateDeck(deck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, deck)
}

// ListDecks lists decks belonging to the authenticated user
func (h *DeckHandler) ListDecks(c *gin.Context) {
	uidAny, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}
	uidStr, ok := uidAny.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID in context"})
		return
	}
	userID, err := uuid.Parse(uidStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID format"})
		return
	}

	decks, err := h.repo.GetDecksByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(decks)
	
	c.JSON(http.StatusOK, decks)
}

// DeleteDeck deletes a deck by ID
func (h *DeckHandler) DeleteDeck(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid deck ID"})
		return
	}
	if err := h.repo.DeleteDeck(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// UpdateDeckTitle updates only the title of a deck
func (h *DeckHandler) UpdateDeckTitle(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid deck ID"})
		return
	}
	var req struct {
		Title string `json:"title" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.UpdateDeckTitle(id, req.Title); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
