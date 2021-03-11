package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*DevuelvoTweetsSeguidores es la estructura con la que devolveremos los tweets*/
type DevuelvoTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID         string             `bson:"usuarioid" json:"usuarioId,omitempty"`
	UsuarioRelacionID string             `bson:"usuariorelacionid" json:"usuarioRelacionId,omitempty"`
	Tweet             struct {
		mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitemtpy"`
	}
}
