package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/kibreab/backend/internal/models"
)

const openAIURL = "https://api.openai.com/v1/completions"

// AIService handles calls to the OpenAI API
type AIService struct {
	apiKey string
}

func NewAIService(apiKey string) *AIService {
	return &AIService{apiKey: apiKey}
}

// GenerateFlashcards calls OpenAI to create Q&A pairs from input text
func (s *AIService) GenerateFlashcards(text string) ([]models.Flashcard, error) {
	payload := map[string]interface{}{
		"model":       "text-davinci-003",
		"prompt":      "Create flashcards (Q: question, A: answer) from this text:\n" + text,
		"max_tokens":  500,
		"temperature": 0.7,
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", openAIURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Choices []struct{ Text string } `json:"choices"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if len(result.Choices) == 0 {
		return nil, errors.New("no flashcards generated")
	}

	return parseFlashcards(result.Choices[0].Text)
}

// parseFlashcards parses raw AI response into []models.Flashcard
func parseFlashcards(raw string) ([]models.Flashcard, error) {
	lines := strings.Split(raw, "\n")
	var cards []models.Flashcard
	for _, line := range lines {
		q := strings.TrimSpace(strings.TrimPrefix(line, "Q:"))
		if strings.HasPrefix(line, "Q:") {
			q = strings.TrimSpace(q)
			// Next line expected A:
			continue
		}
		if strings.HasPrefix(line, "A:") {
			a := strings.TrimSpace(strings.TrimPrefix(line, "A:"))
			cards = append(cards, models.Flashcard{Front: q, Back: a, AIGenerated: true})
		}
	}
	return cards, nil
}
