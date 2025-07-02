package models

import (
	"time"

	"github.com/google/uuid"
)

type Deck struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid;not null;index"`
	Title       string    `gorm:"type:text;not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
