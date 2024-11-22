package repository

import (
	"electomock/internal/domain/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AuthRepository interface {
	CreateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	FindUserByID(id primitive.ObjectID) (*models.User, error)
	UpdateUser(*models.User) error
	FindUserByVerificationCode(code string) (*models.User, error)
	FindUserByProvider(provider string) (*models.User, error)
}

type UserRepository struct {
	db *mongo.Database
}

// FindUserByEmail implements AuthRepository.
func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	panic("unimplemented")
}

// FindUserByID implements AuthRepository.
func (r *UserRepository) FindUserByID(id primitive.ObjectID) (*models.User, error) {
	panic("unimplemented")
}

// FindUserByProvider implements AuthRepository.
func (r *UserRepository) FindUserByProvider(provider string) (*models.User, error) {
	panic("unimplemented")
}

// FindUserByVerificationCode implements AuthRepository.
func (r *UserRepository) FindUserByVerificationCode(code string) (*models.User, error) {
	panic("unimplemented")
}

// UpdateUser implements AuthRepository.
func (r *UserRepository) UpdateUser(*models.User) error {
	panic("unimplemented")
}

func NewAuthRepository(db *mongo.Database) AuthRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) collection() *mongo.Collection {
	return r.db.Collection("users")
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return nil
}
