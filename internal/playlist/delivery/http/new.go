package http

import (
	"github.com/gin-gonic/gin"
	"github.com/quachhoang2002/Music-Library/internal/playlist/usecase"
	pkgLog "github.com/quachhoang2002/Music-Library/pkg/log"
)

type Handler interface {
	List(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Detail(c *gin.Context)
	AddTrack(c *gin.Context)
	RemoveTrack(c *gin.Context)
}

type handler struct {
	l  pkgLog.Logger
	uc usecase.UseCase
}

func New(l pkgLog.Logger, uc usecase.UseCase) Handler {

	return handler{
		l:  l,
		uc: uc,
	}
}
