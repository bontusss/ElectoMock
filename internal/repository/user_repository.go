package repository

import (
	"electomock/internal/domain/models"
	"electomock/internal/infrastructure/database"
)

type UserRepository struct {
	db *database.Database
}

func NewUserRepository(db *database.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r UserRepository) FindByProvider(provider, providerID string) (*models.User, error) {
	var user models.User
	err := r.db.Where("provider = ? AND provider_id = ?", provider, providerID).First(&user).Error
	return &user, err
}
