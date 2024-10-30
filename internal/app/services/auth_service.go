package services

import (
	"electomock/internal/domain/models"
	"electomock/internal/infrastructure/oauth"
	"electomock/internal/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo      *repository.UserRepository
	emailService  *EmailService
	googleOauth   *oauth.GoogleOauth
	facebookOauth *oauth.FacebookOauth
}

func NewAuthService(
	userRepo *repository.UserRepository,
	emailService *EmailService,
	googleOauth *oauth.GoogleOauth,
	facebookOauth *oauth.FacebookOauth) *AuthService {
	return &AuthService{
		userRepo:      userRepo,
		emailService:  emailService,
		googleOauth:   googleOauth,
		facebookOauth: facebookOauth,
	}
}

func (s AuthService) CreateUser(email, password, name string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
		Provider: "local",
	}
	err = s.userRepo.Create(user)
	return user, err
}

func (s AuthService) Login(email, password string) (*models.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
