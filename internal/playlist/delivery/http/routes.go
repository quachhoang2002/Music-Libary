package http

import (
	"github.com/gin-gonic/gin"
	"github.com/quachhoang2002/Music-Library/internal/middleware"
)

func MapMusicPlaylistRoutes(r *gin.RouterGroup, h Handler, mw middleware.Middleware) {
	r.Use(mw.Auth())

	groupRouter := r.Group("/:user_id")

	groupRouter.GET("", h.List)
	groupRouter.POST("", h.Create)
	groupRouter.PUT("/:id", h.Update)
	groupRouter.DELETE("/:id", h.Delete)
	groupRouter.GET("/:id", h.Detail)

	groupRouter.POST("/:id/tracks/:track_id", h.AddTrack)
	groupRouter.DELETE("/:id/tracks/:track_id", h.RemoveTrack)
}
