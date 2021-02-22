package middlew

import (
	"github.com/nelsongp/twittor/routers"
	"net/http"
)

/*ValidoJWT permite validar el jwt que nos viene de la app */
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token !" + err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w,r)
	}
}
