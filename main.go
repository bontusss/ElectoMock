package main

import (
	"context"
	"electomock/config"
	"electomock/internal/app/handlers"
	"electomock/internal/app/middleware"
	"electomock/internal/app/services"
	"electomock/internal/infrastructure/database"
	"electomock/internal/infrastructure/oauth"
	"electomock/internal/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()
	if cfg == nil {
		log.Fatal("Failed to load configuration")
	}

	// Initialize MongoDB connection
	mongo := database.NewMongoDB(cfg.DatabaseUrl)
	db := mongo.Client.Database("electomock")

	// Initialize repositories
	authRepo := repository.NewAuthRepository(db)

	// Initialize services
	emailService := services.NewEmailService(cfg.SMTPConfig)
	googleOAuth := oauth.NewGoogleOAuth(cfg.GoogleConfig)
	facebookOAuth := oauth.NewfacebookOAuth(cfg.FacebookConfig)
	authService := services.NewAuthService(authRepo, *emailService, cfg)

	// Initialize handlers
	authHandler := handlers.NewAuthService(authService)

	// Initialize Gin router
	router := gin.Default()

	// Configure middleware
	router.Use(middleware.Sessions(cfg.SessionSecret))

	// Load templates and custom functions
	router.LoadHTMLGlob("templates/**/*.html")

	// Serve static files
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Home",
		})
	})

	router.GET("/about", func(c *gin.Context) {
		c.HTML(200, "about.html", gin.H{
			"title": "About",
		})
	})
	router.GET("/privacy-policy", func(c *gin.Context) {
		c.HTML(200, "public/privacy-policy.html", gin.H{})
	})

	// Configure routes
	auth := router.Group("/auth")
	{
		auth.GET("/login", func(c *gin.Context) {
			c.HTML(200, "auth/login.html", gin.H{})
		})
		auth.GET("/signup", func(c *gin.Context) {
			c.HTML(200, "auth/signup.html", gin.H{})
		})
		auth.POST("/register", authHandler.Register)

		// OAuth routes
		auth.GET("/google", func(c *gin.Context) {
			c.Redirect(302, googleOAuth.GetAuthURL())
		})
		auth.GET("/facebook", func(c *gin.Context) {
			c.Redirect(302, facebookOAuth.GetAuthURL())
		})
	}

	// Protected routes
	protected := router.Group("/mock")
	protected.Use(middleware.AuthRequired())
	{
		// protected.GET("/", func(c *gin.Context) {
		// 	c.HTML(200, "layout/base.html", gin.H{})
		// })
	}

	// Create server with graceful shutdown
	srv := &http.Server{
		Addr:    ":4000",
		Handler: router,
	}

	// Start server in goroutine to not block
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Shutdown server with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
