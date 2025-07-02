package models

import (
	"time"

	"github.com/google/uuid"
)

type Flashcard struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	DeckID      uuid.UUID `gorm:"type:uuid;not null;index"`
	Front       string    `gorm:"type:text;not null"`
	Back        string    `gorm:"type:text;not null"`
	AIGenerated bool      `gorm:"type:boolean;default:false"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
