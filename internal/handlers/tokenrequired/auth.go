package tokenrequired

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sembh1998/wedding_invitation_generation/internal/core/domain/token"
)

func TokenRequiredMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the value of the "my_cookie" cookie
		cookieValue, err := c.Cookie("wedding_cookie")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "cookie not found"})
			return
		}

		payload, err := token.NewTokenizer().Maker.VerifyToken(cookieValue)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("user_id", payload.UserID)
		c.Set("username", payload.Username)

		c.Next()

	}
}

func BearerTokenRequiredMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header not found"})
			return
		}

		// Expected format: "Bearer <token>"
		const prefix = "Bearer "
		if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			return
		}

		tokenStr := authHeader[len(prefix):]

		payload, err := token.NewTokenizer().Maker.VerifyToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("user_id", payload.UserID)
		c.Set("username", payload.Username)

		c.Next()
	}
}
