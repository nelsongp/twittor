package middlew

import (
	"github.com/nelsongp/twittor/bd"
	"net/http"
)

/*ChequeoBD el curso dice que para que sea un middleware debe enviar lo que recibe es decir
debe ser un pasamano en este caso recibe un http y regresa un http*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		//se envia un http error custom con el codigo 500 si la conexion eta mala
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la Base de datos", 500)
			//al retornar vacio es como que ahi termina todo
			return
		}
		next.ServeHTTP(w,r)
	}
}
