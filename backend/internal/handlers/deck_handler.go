package handlers

import (
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

// CreateDeck handles POST /api/decks
func (h *DeckHandler) CreateDeck(c *gin.Context) {
    var req struct {
        UserID      uuid.UUID `json:"user_id" binding:"required"`
        Title       string    `json:"title" binding:"required"`
        Description string    `json:"description"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    deck := &models.Deck{
        UserID:      req.UserID,
        Title:       req.Title,
        Description: req.Description,
    }

    if err := h.repo.CreateDeck(deck); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, deck)
}

// ListDecks handles GET /api/decks?user_id=...
func (h *DeckHandler) ListDecks(c *gin.Context) {
    userIDStr := c.Query("user_id")
    userID, err := uuid.Parse(userIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
        return
    }

    decks, err := h.repo.GetDecksByUser(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, decks)
}

// DeleteDeck handles DELETE /api/decks/:id
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

// UpdateDeckTitle handles PATCH /api/decks/:id
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
