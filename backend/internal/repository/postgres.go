package repository

import (
	"github.com/kibreab/backend/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresRepo struct {
	DB *gorm.DB
}

type Repository interface {
	AutoMigrate() error

	// User methods
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)

	// Deck methods
	CreateDeck(deck *models.Deck) error
	GetDecksByUser(userID uuid.UUID) ([]models.Deck, error)
	GetDeckByID(deckID uuid.UUID) (*models.Deck, error)
	DeleteDeck(deckID uuid.UUID) error
	UpdateDeckTitle(deckID uuid.UUID, newTitle string) error

	// Flashcard methods
	CreateFlashcard(card *models.Flashcard) error
	GetFlashcardsByDeck(deckID uuid.UUID) ([]models.Flashcard, error)
	GetFlashcardByID(id uuid.UUID) (*models.Flashcard, error)
	DeleteFlashcard(id uuid.UUID) error
}

func NewPostgresRepo(db *gorm.DB) *PostgresRepo {
	return &PostgresRepo{DB: db}
}

func (r *PostgresRepo) AutoMigrate() error {
	return r.DB.
		AutoMigrate(
			&models.User{},
			&models.Deck{},
			&models.Flashcard{},
		)
}
