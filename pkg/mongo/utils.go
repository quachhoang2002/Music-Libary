package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ObjectIDFromHexOrNil returns an ObjectID from the provided hex representation.
func ObjectIDFromHexOrNil(id string) primitive.ObjectID {
	objID, _ := primitive.ObjectIDFromHex(id)
	return objID
}

func BuildQueryWithSoftDelete(query bson.M) bson.M {
	query["deleted_at"] = nil
	return query
}
