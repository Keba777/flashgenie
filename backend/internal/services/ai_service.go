package services

import (
    "bytes"
    
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"

    "github.com/kibreab/backend/internal/models"
)

const togetherAIURL = "https://api.together.xyz/v1/chat/completions"

// AIService handles calls to the Together AI API
type AIService struct {
    apiKey string
}

// NewAIService initializes the AI service
func NewAIService(apiKey string) *AIService {
    return &AIService{
        apiKey: apiKey,
    }
}

// GenerateFlashcards generates flashcards from the provided text using Together API
func (s *AIService) GenerateFlashcards(text string) ([]models.Flashcard, error) {
    if s.apiKey == "" {
        return nil, errors.New("together API key is not set")
    }

    prompt := fmt.Sprintf(`Create flashcards from the following text and output them as a JSON object with a "flashcards" field containing an array of objects, each with "question" and "answer" fields.\n\nText:\n%s`, text)

    payload := map[string]interface{}{
        "model": "mistralai/Mixtral-8x7B-Instruct-v0.1",
        "messages": []map[string]string{
            {
                "role":    "user",
                "content": prompt,
            },
        },
        "temperature": 0.7,
        "max_tokens": 500,
    }

    bodyBytes, err := json.Marshal(payload)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal payload: %w", err)
    }

    req, err := http.NewRequest("POST", togetherAIURL, bytes.NewReader(bodyBytes))
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }
    req.Header.Set("Authorization", "Bearer "+s.apiKey)
    req.Header.Set("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("request error: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        return nil, fmt.Errorf("together API error: status %d, message: %s", resp.StatusCode, string(body))
    }

    var result struct {
        Choices []struct {
            Message struct {
                Content string `json:"content"`
            } `json:"message"`
        } `json:"choices"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("failed to decode response: %w", err)
    }

    if len(result.Choices) == 0 {
        return nil, errors.New("no flashcards generated")
    }

    rawText := result.Choices[0].Message.Content
    fmt.Println("Together AI raw response text:\n", rawText)

    var parsed struct {
        Flashcards []struct {
            Question string `json:"question"`
            Answer   string `json:"answer"`
        } `json:"flashcards"`
    }

    if err := json.Unmarshal([]byte(rawText), &parsed); err != nil {
        return nil, fmt.Errorf("failed to parse flashcard JSON: %w", err)
    }

    var cards []models.Flashcard
    for _, fc := range parsed.Flashcards {
        cards = append(cards, models.Flashcard{
            Front:       fc.Question,
            Back:        fc.Answer,
            AIGenerated: true,
        })
    }

    return cards, nil
}

// ListModels returns a placeholder since Together doesn't support listing models via API
func (s *AIService) ListModels() ([]string, error) {
    return []string{"mistralai/Mixtral-8x7B-Instruct-v0.1"}, nil
}
