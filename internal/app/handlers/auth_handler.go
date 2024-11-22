package handlers

import (
	"electomock/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthService(service services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (a *AuthHandler) Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "auth/register.html", gin.H{
			"error": "Invalid input, Please chheck your details.",
		})
		return
	}
	if err := a.service.Register(input.Name, input.Email, input.Password); err != nil {
		c.HTML(http.StatusBadRequest, "auth/register.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusOK, "/")
}
