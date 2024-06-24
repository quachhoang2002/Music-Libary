package http

import (
	"github.com/gin-gonic/gin"
	"github.com/quachhoang2002/Music-Library/internal/middleware"
)

func MapMusicTrackRoutes(r *gin.RouterGroup, h Handler, mw middleware.Middleware) {
	r.Use(mw.Auth())

	r.GET("", h.List)
	r.GET("/:id", h.Detail)
	r.POST("", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id/file", h.GetFile)
}
