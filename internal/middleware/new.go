package middleware

import (
	"github.com/xuanhoang/music-library/pkg/jwt"
	"github.com/xuanhoang/music-library/pkg/log"
)

type Middleware struct {
	l          log.Logger
	jwtManager jwt.Maker
}

func New(l log.Logger, jwtManager jwt.Maker) Middleware {
	return Middleware{
		l:          l,
		jwtManager: jwtManager,
	}
}
