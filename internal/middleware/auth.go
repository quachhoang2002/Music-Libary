package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/quachhoang2002/Music-Library/pkg/jwt"
	"github.com/quachhoang2002/Music-Library/pkg/response"
)

var ()

func (m Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", "")
		if tokenString == "" {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		payload, err := m.jwtManager.VerifyToken(tokenString)
		if err != nil {
			// response.Unauthorized(c)
			// c.Abort()
			// return
		}

		ctx := c.Request.Context()
		ctx = jwt.SetPayloadToContext(ctx, payload)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
