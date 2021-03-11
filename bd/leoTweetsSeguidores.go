package bd

import (
	"context"
	"github.com/nelsongp/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

/*LeoTweetsSeguidores leo los tweets de mis seguidores*/
func LeoTweetsSeguidores (ID string, pagina int)([]models.DevuelvoTweetsSeguidores, bool){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	skip := (pagina -1) * 20

	condiciones := make([]bson.M,0)
	//este seria el usurio con el que haria match
	condiciones = append(condiciones, bson.M{"$match":bson.M{"usuarioid":ID}})
	//ahora se construye la condicion
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	//lo que hace el unwind es traerte ordenado los objetos, si no estuviera lo muestra como un maestro detalle
	condiciones = append(condiciones, bson.M{"$unwind":"$tweet"})
	condiciones = append(condiciones, bson.M{"$sort":bson.M{"tweet.fecha":-1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit":20})

	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.DevuelvoTweetsSeguidores
	err = cursor.All(ctx, &result)

	if err != nil {
		return result, false
	}

	return result, true
}
