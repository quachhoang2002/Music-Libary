package mongo

import (
	"context"

	"github.com/xuanhoang/music-library/internal/models"
	"github.com/xuanhoang/music-library/internal/music/repository"
	"github.com/xuanhoang/music-library/pkg/mongo"
	"github.com/xuanhoang/music-library/pkg/paginator"
	"github.com/xuanhoang/music-library/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	musicTrackCollection = "music_tracks"
)

func (r implRepository) getTrackCollection() mongo.Collection {
	return r.database.Collection(musicTrackCollection)
}
func (repo implRepository) Create(ctx context.Context, sc models.Scope, opt repository.CreateOpt) (models.MusicTrack, error) {
	c := repo.getTrackCollection()

	track := models.MusicTrack{
		ID:            primitive.NewObjectID(),
		Title:         opt.Title,
		Artist:        opt.Artist,
		Album:         opt.Album,
		Genre:         opt.Genre,
		ReleaseYear:   opt.ReleaseYear,
		Duration:      opt.Duration,
		MP3FilePath:   opt.MP3FilePath,
		CreatedUserID: mongo.ObjectIDFromHexOrNil(sc.UserID),
		CreatedAt:     util.Now(),
		UpdatedAt:     util.Now(),
	}

	_, err := c.InsertOne(ctx, track)
	if err != nil {
		repo.l.Error(ctx, "music.repository.Create.InsertOne", err.Error())
		return models.MusicTrack{}, err
	}

	return track, nil
}

func (r implRepository) Update(ctx context.Context, sc models.Scope, opt repository.UpdateOpt) (models.MusicTrack, error) {
	c := r.getTrackCollection()

	filter := mongo.BuildQueryWithSoftDelete(bson.M{})
	filter["_id"] = mongo.ObjectIDFromHexOrNil(opt.ID)

	track := models.MusicTrack{
		Title:       opt.Data.Title,
		Artist:      opt.Data.Artist,
		Album:       opt.Data.Album,
		Genre:       opt.Data.Genre,
		ReleaseYear: opt.Data.ReleaseYear,
		Duration:    opt.Data.Duration,
		UpdatedAt:   util.Now(),
		MP3FilePath: opt.Data.MP3FilePath,
	}

	_, err := c.UpdateOne(ctx, filter, bson.M{"$set": track})
	if err != nil {
		r.l.Error(ctx, "music.repository.Create.InsertOne", err.Error())
		return models.MusicTrack{}, err
	}

	return track, nil
}

func (r implRepository) List(ctx context.Context, sc models.Scope, opt repository.ListOpt) ([]models.MusicTrack, paginator.Paginator, error) {
	c := r.getTrackCollection()

	filterQuery := mongo.BuildQueryWithSoftDelete(bson.M{})
	filter := opt.Filter
	if filter.Title != "" {
		//query where like
		filterQuery["title"] = primitive.Regex{Pattern: filter.Title, Options: "i"}
	}
	if filter.Artist != "" {
		//query where like
		filterQuery["artist"] = primitive.Regex{Pattern: filter.Artist, Options: "i"}
	}
	if filter.Album != "" {
		filterQuery["album"] = filter.Album
	}

	if len(filter.IDs) > 0 {
		filterQuery["_id"] = bson.M{"$in": mongo.ObjectIDsFromHex(filter.IDs)}
	}

	cur, err := c.Find(ctx, filterQuery, options.Find().
		SetSkip(opt.PaginatorQuery.Offset()).
		SetLimit(opt.PaginatorQuery.Limit).
		SetSort(bson.D{
			{Key: "created_at", Value: -1},
			{Key: "_id", Value: -1},
		}))
	if err != nil {
		r.l.Error(ctx, "music.repository.Get.Find", err.Error())
		return nil, paginator.Paginator{}, err
	}

	var tracks []models.MusicTrack
	err = cur.All(ctx, &tracks)
	if err != nil {
		r.l.Error(ctx, "music.repository.Get.All", err.Error())
		return nil, paginator.Paginator{}, err
	}

	total, err := c.CountDocuments(ctx, filterQuery)
	if err != nil {
		r.l.Error(ctx, "music.repository.Get.CountDocuments", err.Error())
		return nil, paginator.Paginator{}, err
	}

	return tracks, paginator.Paginator{
		Total:       total,
		Count:       int64(len(tracks)),
		PerPage:     opt.PaginatorQuery.Limit,
		CurrentPage: opt.PaginatorQuery.Page,
	}, nil
}

func (r implRepository) Delete(ctx context.Context, sc models.Scope, id string) error {
	c := r.getTrackCollection()

	filter := mongo.BuildQueryWithSoftDelete(bson.M{})
	filter["_id"] = mongo.ObjectIDFromHexOrNil(id)

	_, err := c.DeleteOneSoft(ctx, filter)
	if err != nil {
		r.l.Error(ctx, "music.repository.Delete.UpdateOne", err.Error())
		return err
	}

	return nil
}

func (r implRepository) Detail(ctx context.Context, sc models.Scope, id string) (models.MusicTrack, error) {
	c := r.getTrackCollection()

	filter := mongo.BuildQueryWithSoftDelete(bson.M{})
	filter["_id"] = mongo.ObjectIDFromHexOrNil(id)

	var track models.MusicTrack
	err := c.FindOne(ctx, filter).Decode(&track)
	if err != nil {
		r.l.Error(ctx, "music.repository.Detail.FindOne", err.Error())
		return models.MusicTrack{}, err
	}

	return track, nil
}
