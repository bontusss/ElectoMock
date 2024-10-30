package handlers

import (
	"electomock/internal/app/services"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthhandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}
func (h AuthHandler) HomeHandler(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{"app_name": "ElectoMock"})
}

func (h AuthHandler) LoginPage(c *gin.Context) {
	c.HTML(200, "auth/login.html", gin.H{"title": "login"})
}

func (h AuthHandler) Login(c *gin.Context) {
	var form struct {
		Email    string `form:"email" binding:"required,email"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&form); err != nil {
		c.HTML(400, "auth/login.html", gin.H{
			"error": "invalid form data.",
		})
		return
	}
	user, err := h.authService.Login(form.Email, form.Password)
	if err != nil {
		c.HTML(401, "auth/login.html", gin.H{
			"error": "invalid credentials.",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	err = session.Save()
	if err != nil {
		c.HTML(500, "auth/login.html", gin.H{
			"error": "An error occurred, try again.",
		})
		return
	}

	// todo: redirect to homepage, change it to user dashboard when implemented.
	c.Redirect(302, "/")
}

func (h AuthHandler) GoogleLogin(c *gin.Context) {
	url := h.authService.
}