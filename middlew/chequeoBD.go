package middlew

import (
	"net/http"

	"github.com/nelsongp/twittor/bd"
)

/*ChequeoBD es el middleware que me permite conocer el estado de laBF*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConexion() == 0 {
			//con el 500 al final es el tipo de error que voy a retornar
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
