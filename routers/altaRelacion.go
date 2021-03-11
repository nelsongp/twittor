package routers

import (
	"github.com/nelsongp/twittor/bd"
	"github.com/nelsongp/twittor/models"
	"net/http"
)

/*AltaRelacion realiza el regitro de la relacion entre usuarios*/
func AltaRelacion(w http.ResponseWriter, r* http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un problema al insertar la relacion", http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar la relacion", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
