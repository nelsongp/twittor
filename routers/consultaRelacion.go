package routers

import (
	"encoding/json"
	"github.com/nelsongp/twittor/bd"
	"github.com/nelsongp/twittor/models"
	"net/http"
)

/*ConsultaRelacion consulta el regitro de la relacion entre usuarios*/
func ConsultaRelacion(w http.ResponseWriter, r* http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
