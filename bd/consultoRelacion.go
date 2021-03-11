package bd

import (
	"context"
	"fmt"
	"github.com/nelsongp/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

/*ConsultoRelacion consulta relacion de usuarios*/
func ConsultoRelacion(t models.Relacion)(bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid": t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		return false, err
	}
	return true, nil
}
