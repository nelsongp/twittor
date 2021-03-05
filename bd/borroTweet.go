package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*Borro tweet borra un tweet determinado*/
func BorroTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(ctx, condicion)

	return err
}
