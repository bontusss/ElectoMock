package services

import (
	"electomock/config"
	"electomock/internal/domain/models"
	"electomock/internal/repository"
)

type AuthService interface {
	Register(name, email, password string) error
	Login(email, password string) (*models.User, error)
	VerifyEmail(code string) error
	RequestPasswordreset(email string) error
	ResetPassword(token, newPassword string) error
	GoogleCallback(code string) (*models.User, error)
	FacebookCall(code string) (*models.User, error)
}

type authService struct {
	repo      repository.AuthRepository
	mailer    EmailService
	appConfig *config.Config
}

// FacebookCall implements AuthService.
func (a *authService) FacebookCall(code string) (*models.User, error) {
	panic("unimplemented")
}

// GoogleCallback implements AuthService.
func (a *authService) GoogleCallback(code string) (*models.User, error) {
	panic("unimplemented")
}

// Login implements AuthService.
func (a *authService) Login(email string, password string) (*models.User, error) {
	panic("unimplemented")
}

// Register implements AuthService.
func (a *authService) Register(name string, email string, password string) error {
	panic("unimplemented")
}

// RequestPasswordreset implements AuthService.
func (a *authService) RequestPasswordreset(email string) error {
	panic("unimplemented")
}

// ResetPassword implements AuthService.
func (a *authService) ResetPassword(token string, newPassword string) error {
	panic("unimplemented")
}

// VerifyEmail implements AuthService.
func (a *authService) VerifyEmail(code string) error {
	panic("unimplemented")
}

func NewAuthService(repo repository.AuthRepository, mailer EmailService, appConfig *config.Config) AuthService {
	return &authService{repo: repo, mailer: mailer, appConfig: appConfig}
}
