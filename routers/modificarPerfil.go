package routers

import (
	"encoding/json"
	"github.com/nelsongp/twittor/bd"
	"github.com/nelsongp/twittor/models"
	"net/http"
)

/*Modificar es la funcion para hacer update de un usuario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentra modificar el registro. Reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se pudo modificar el registro. Reintente nuevamente"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
