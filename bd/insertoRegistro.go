package bd

import (
	"context"
	"github.com/nelsongp/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*InsertoRegistro es la parada final coin la BD para insertar los datos del usuario */
func InsertoRegistro (u models.Usuario)(string, bool, error){

	//el contexto de background es el que seguimos jalando desde el de base de datos
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	/* Se hace el timeout para que cierre el contexto luego de ejecutar una orden,
	ya que si no lo realizo se empiezan a acumular diferentetes peticiones dentro del
	context, ya que este es uno en general para todo el proyecto */

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	//la funcion insertedID retuorna el id del objeto insertado el ObjectID del model usurio
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
