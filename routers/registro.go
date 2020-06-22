package routers

import (
	"encoding/json"
	"net/http"

	"github.com/nelsongp/twittor/bd"
	"github.com/nelsongp/twittor/models"
)

/*Registro es la fnucion para crear en la BD el registro del usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	//cuando se ocupa el r.Body solo se puede ocupar una sola vez, cuando se lee por primera vez se destruye y es de tipo lecura
	//le decimos que me lo guarde en el puntero de models de usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contrasenia de al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logro insertar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
