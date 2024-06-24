package middleware

import (
	"github.com/quachhoang2002/Music-Library/pkg/jwt"
	"github.com/quachhoang2002/Music-Library/pkg/log"
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
