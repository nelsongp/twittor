package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion a la base de datos*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://twittor:twittor@cluster0-ruzay.mongodb.net/<dbname>?retryWrites=true&w=majority")

/*ConectarBD es la funcion que me permite conecatar con la BD*/
func ConectarBD() *mongo.Client {
	//los contextos son espacios en memoras que sirven para compartir cosas, para comunicar informacion entre
	//funciones y evitar los timeouts
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	//validar si la base de datos conecta de manera correcta
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion Exitosa con la BD")
	return client
}

/*ChequeoConexion es el ping a la base de datos*/
func ChequeoConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
