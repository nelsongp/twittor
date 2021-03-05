package routers

import (
	"github.com/nelsongp/twittor/bd"
	"github.com/nelsongp/twittor/models"
	"io"
	"net/http"
	"os"
	"strings"
)

/*SubirBanner sube el banners a un usuario*/
func SubirBanner(w http.ResponseWriter, r *http.Request){
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	//el 0666 son los permisos para poder modificar los archivos
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir la imagen ! " + err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error al copiar la imagen ! " + err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error al guardar el banner en la BD ! " + err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application-json")
	w.WriteHeader(http.StatusCreated)
}
