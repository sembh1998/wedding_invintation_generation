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
