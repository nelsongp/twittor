package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DevuelvoTweets struct {
	ID      primitive.ObjectID `bson: "_id" json:"_id,ommiyempty"`
	UserID  string             `bson: "userid" json:"userID,ommiyempty"`
	Mensaje string             `bson: "mensaje" json:"mensaje,ommiyempty"`
	Fecha   time.Time          `bson: "fecha" json:"fecha,ommiyempty"`
}
