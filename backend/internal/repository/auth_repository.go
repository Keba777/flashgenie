package repository

import (
    "errors"

    "gorm.io/gorm"
    "github.com/kibreab/backend/internal/models"
)

// CreateUser inserts a new User record.
func (r *PostgresRepo) CreateUser(user *models.User) error {
    return r.DB.Create(user).Error
}

// GetUserByEmail fetches a user by their email.
func (r *PostgresRepo) GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("user not found")
        }
        return nil, err
    }
    return &user, nil
}
