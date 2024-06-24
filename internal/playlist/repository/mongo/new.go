package mongo

import (
	"github.com/quachhoang2002/Music-Library/internal/playlist/repository"
	pkgLog "github.com/quachhoang2002/Music-Library/pkg/log"
	"github.com/quachhoang2002/Music-Library/pkg/mongo"
)

type implRepository struct {
	l        pkgLog.Logger
	database mongo.Database
}

var _ repository.Repository = implRepository{}

func New(l pkgLog.Logger, database mongo.Database) implRepository {
	return implRepository{
		l:        l,
		database: database,
	}
}
