package mongo

import (
	"github.com/xuanhoang/music-library/internal/music/repository"
	pkgLog "github.com/xuanhoang/music-library/pkg/log"
	"github.com/xuanhoang/music-library/pkg/mongo"
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
