package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/quachhoang2002/Music-Library/pkg/response"
	"github.com/quachhoang2002/Music-Library/pkg/telegram"
)

func Recovery(t telegram.Telegram, chatBugID int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[Panic Recovered] %v\n", err)
				log.Printf("Request: %s %s\n", c.Request.Method, c.Request.URL.Path)
				response.PanicError(c, err, t, chatBugID)
			}
		}()
		c.Next()
	}
}
