package main

import (
	"electomock/config"
	"electomock/internal/app/services"
	"electomock/internal/infrastructure/database"
	"electomock/internal/infrastructure/oauth"
	"electomock/internal/repository"
)

func main() {
	cfg := config.Load()

	//	Initialize DB
	db := database.NewPostgresDB(cfg.DatabaseUrl)
	defer db.Close()

	//	Initialize Repositories
	userRepo := repository.NewUserRepository(db)

	//	Initialize Services
	emailService := services.NewEmailService(cfg.SMTPConfig)
	authService := services.NewAuthService(
		userRepo,
		emailService,
		oauth.NewGoogleOAuth(cfg.GoogleConfig),
		oauth.NewfacebookOAuth(cfg.FacebookConfig),
	)
}
