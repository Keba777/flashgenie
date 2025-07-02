package repository

import (
	"github.com/google/uuid"

	"github.com/kibreab/backend/internal/models"
)

func (r *PostgresRepo) CreateDeck(deck *models.Deck) error {
	return r.DB.Create(deck).Error
}

func (r *PostgresRepo) GetDeckByID(deckID uuid.UUID) (*models.Deck, error) {
	var deck models.Deck
	if err := r.DB.First(&deck, "id = ?", deckID).Error; err != nil {
		return nil, err
	}
	return &deck, nil
}

func (r *PostgresRepo) GetDecksByUser(userID uuid.UUID) ([]models.Deck, error) {
	var decks []models.Deck
	if err := r.DB.Where("user_id = ?", userID).Find(&decks).Error; err != nil {
		return nil, err
	}
	return decks, nil
}

func (r *PostgresRepo) DeleteDeck(deckID uuid.UUID) error {
	return r.DB.Delete(&models.Deck{}, "id = ?", deckID).Error
}

func (r *PostgresRepo) UpdateDeckTitle(deckID uuid.UUID, newTitle string) error {
	return r.DB.Model(&models.Deck{}).
		Where("id = ?", deckID).
		Update("title", newTitle).Error
}
