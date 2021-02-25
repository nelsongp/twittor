package routers

import (
	"encoding/json"
	"github.com/nelsongp/twittor/bd"
	"github.com/nelsongp/twittor/models"
	"net/http"
	"time"
)

/*GraboTweet perimte grabar el twweet en la base de dats*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se pudo crear el tweet"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
