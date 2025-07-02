package repository

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/kibreab/backend/internal/models"
)

// CreateFlashcard inserts a new flashcard into the database.
func (r *PostgresRepo) CreateFlashcard(card *models.Flashcard) error {
	return r.DB.Create(card).Error
}

// GetFlashcardsByDeck retrieves all flashcards associated with a specific deck.
func (r *PostgresRepo) GetFlashcardsByDeck(deckID uuid.UUID) ([]models.Flashcard, error) {
	var cards []models.Flashcard
	if err := r.DB.Where("deck_id = ?", deckID).Find(&cards).Error; err != nil {
		return nil, err
	}
	return cards, nil
}

// Optional: GetFlashcardByID retrieves a single flashcard by ID.
func (r *PostgresRepo) GetFlashcardByID(id uuid.UUID) (*models.Flashcard, error) {
	var card models.Flashcard
	if err := r.DB.First(&card, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("flashcard not found")
		}
		return nil, err
	}
	return &card, nil
}

// Optional: DeleteFlashcard removes a flashcard from the DB.
func (r *PostgresRepo) DeleteFlashcard(id uuid.UUID) error {
	return r.DB.Delete(&models.Flashcard{}, "id = ?", id).Error
}
