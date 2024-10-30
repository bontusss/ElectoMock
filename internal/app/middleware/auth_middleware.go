package middleware

import (
	"electomock/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")
		if userID == nil {
			c.Redirect(302, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}

func Sessions(secret string) gin.HandlerFunc {
	cfg := config.Config{}
	store := cookie.NewStore([]byte(cfg.SessionSecret))
	store.Options(sessions.Options{
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		MaxAge:   86300 * 64,
	})
	return sessions.Sessions("session", store)
}
