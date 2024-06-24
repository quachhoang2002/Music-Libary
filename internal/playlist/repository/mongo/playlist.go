package mongo

import (
	"context"

	"github.com/quachhoang2002/Music-Library/internal/models"
	"github.com/quachhoang2002/Music-Library/internal/playlist/repository"
	"github.com/quachhoang2002/Music-Library/pkg/mongo"
	"github.com/quachhoang2002/Music-Library/pkg/paginator"
	"github.com/quachhoang2002/Music-Library/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	musicTrackCollection = "playlists"
)

func (r implRepository) getCollection() mongo.Collection {
	return r.database.Collection(musicTrackCollection)
}
func (repo implRepository) Create(ctx context.Context, sc models.Scope, opt repository.CreateOpt) (models.Playlist, error) {
	c := repo.getCollection()

	track := models.Playlist{
		ID:        primitive.NewObjectID(),
		Name:      opt.Name,
		UserID:    opt.UserID,
		TrackIDs:  []primitive.ObjectID{},
		CreatedAt: util.Now(),
		UpdatedAt: util.Now(),
	}

	_, err := c.InsertOne(ctx, track)
	if err != nil {
		repo.l.Error(ctx, "music.repository.Create.InsertOne", err.Error())
		return models.Playlist{}, err
	}

	return track, nil
}

func (r implRepository) Update(ctx context.Context, sc models.Scope, opt repository.UpdateOpt) error {
	c := r.getCollection()

	filter := mongo.BuildQueryWithSoftDelete(bson.M{})
	filter["_id"] = mongo.ObjectIDFromHexOrNil(opt.ID)
	filter["user_id"] = opt.UserID

	update := bson.M{}
	if opt.Data.Name != "" {
		update["name"] = opt.Data.Name
	}

	_, err := c.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		r.l.Error(ctx, "music.repository.Create.InsertOne", err.Error())
		return err
	}

	return nil
}

func (r implRepository) AddTrack(ctx context.Context, sc models.Scope, playlistID string, trackID string) error {
	c := r.getCollection()

	filter := mongo.BuildQueryWithSoftDelete(bson.M{})
	filter["_id"] = mongo.ObjectIDFromHexOrNil(playlistID)

	update := bson.M{"$addToSet": bson.M{"track_ids": mongo.ObjectIDFromHexOrNil(trackID)}}
	_, err := c.UpdateOne(ctx, filter, update)
	if err != nil {
		r.l.Error(ctx, "music.repository.AddTrack.UpdateOne", err.Error())
		return err
	}

	return nil
}

// remove track
func (r implRepository) RemoveTrack(ctx context.Context, sc models.Scope, playlistID string, trackID string) error {
	c := r.getCollection()
	filter := mongo.BuildQueryWithSoftDelete(bson.M{})
	filter["_id"] = mongo.ObjectIDFromHexOrNil(playlistID)

	update := bson.M{"$pull": bson.M{"track_ids": mongo.ObjectIDFromHexOrNil(trackID)}}
	_, err := c.UpdateOne(ctx, filter, update)
	if err != nil {
		r.l.Error(ctx, "music.repository.RemoveTrack.UpdateOne", err.Error())
		return err
	}

	return nil
}

func (r implRepository) List(ctx context.Context, sc models.Scope, opt repository.ListOpt) ([]models.Playlist, paginator.Paginator, error) {
	c := r.getCollection()
	filterQuery := mongo.BuildQueryWithSoftDelete(bson.M{})
	filterQuery["user_id"] = opt.UserID

	if len(opt.Filter.TrackIDs) > 0 {
		filterQuery["track_ids"] = bson.M{"$in": mongo.ObjectIDsFromHex(opt.Filter.TrackIDs)}
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

	var playlists []models.Playlist
	err = cur.All(ctx, &playlists)
	if err != nil {
		r.l.Error(ctx, "music.repository.Get.All", err.Error())
		return nil, paginator.Paginator{}, err
	}

	total, err := c.CountDocuments(ctx, filterQuery)
	if err != nil {
		r.l.Error(ctx, "music.repository.Get.CountDocuments", err.Error())
		return nil, paginator.Paginator{}, err
	}

	return playlists, paginator.Paginator{
		Total:       total,
		Count:       int64(len(playlists)),
		PerPage:     opt.PaginatorQuery.Limit,
		CurrentPage: opt.PaginatorQuery.Page,
	}, nil
}

func (r implRepository) Delete(ctx context.Context, sc models.Scope, id string) error {
	c := r.getCollection()

	filter := mongo.BuildQueryWithSoftDelete(bson.M{})
	filter["_id"] = mongo.ObjectIDFromHexOrNil(id)

	_, err := c.DeleteOneSoft(ctx, filter)
	if err != nil {
		r.l.Error(ctx, "music.repository.Delete.UpdateOne", err.Error())
		return err
	}

	return nil
}

func (r implRepository) Detail(ctx context.Context, sc models.Scope, id string) (models.Playlist, error) {
	c := r.getCollection()

	filter := mongo.BuildQueryWithSoftDelete(bson.M{})
	filter["_id"] = mongo.ObjectIDFromHexOrNil(id)

	var playlist models.Playlist
	err := c.FindOne(ctx, filter).Decode(&playlist)
	if err != nil {
		r.l.Error(ctx, "music.repository.Detail.FindOne", err.Error())
		return models.Playlist{}, err
	}

	return playlist, nil
}
