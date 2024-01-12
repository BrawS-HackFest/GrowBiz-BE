package middleware

import (
	"HackFest/config"
	"HackFest/utils"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		client := config.FirebaseClient()
		auth := c.Request.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.Abort()
			utils.HttpInternalError(c, "wrong header value", errors.New("wrong header value"))
			return
		}

		token := auth[7:]
		data, err := client.VerifyIDToken(context.Background(), token)
		if err != nil {
			c.Abort()
			utils.HttpFailOrError(c, 401, "Unauthorized", err)
			return
		}

		c.Set("userID", data.UID)
		c.Next()
	}
}
