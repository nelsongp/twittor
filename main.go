package main

import (
	"github.com/nelsongp/twittor/bd"
	"github.com/nelsongp/twittor/handlers"
	"log"
)

func main(){
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
