package bd

import (
	"context"
	"github.com/nelsongp/twittor/models"
	"time"
)

/*BorroRelacion borra relacion de usuarios*/
func BorroRelacion(t models.Relacion)(bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
