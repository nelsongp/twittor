package routers

import (
	"encoding/json"
	"github.com/nelsongp/twittor/bd"
	"github.com/nelsongp/twittor/models"
	"net/http"
)

/*Registro es la funcion para crear usuario en base de datos*/
func Registro(w http.ResponseWriter, r *http.Request){
	var t models.Usuario
	//leemos el body del request, pero el body es un objeto de solo lectura que se destruye una vez leido
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos definidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0{
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Error contrasenia debe ser mayor a 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil{
		http.Error(w, "Ocurrio un error al registrar usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha podido insertar el registro a la base", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
