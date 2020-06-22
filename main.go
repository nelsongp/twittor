package main

import (
	"log"

	"github.com/nelsongp/twittor/bd"
	"github.com/nelsongp/twittor/handlers"
)

func main() {
	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
