package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kibreab/backend/internal/services"
)

type TestHandler struct {
	aiService *services.AIService
}

func NewTestHandler(aiService *services.AIService) *TestHandler {
	return &TestHandler{aiService: aiService}
}

func (h *TestHandler) CheckOpenAIKey(c *gin.Context) {
	models, err := h.aiService.ListModels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "OpenAI key is invalid or request failed",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":                "OpenAI API key is valid ✅",
		"available_models_count": len(models),
	})
}
