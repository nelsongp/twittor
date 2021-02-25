package bd

import (
	"context"
	"github.com/nelsongp/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)
/*LeoTweets lee los tweets de un perfil especifo*/
func LeoTweets(ID string, pagina int64)([]*models.DevuelvoTweets, bool){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key:"fecha", Value: -1}})
	//es la acumulacion de paginas que se iran al pasar las paginas
	opciones.SetSkip((pagina -1)*20)

	cursor, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	//el contexttodo es para no utilizar le crado de la base por 15 seg y usar uno propio temporal
	for cursor.Next(context.TODO()){
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados,&registro)
	}
	return resultados, true
}
