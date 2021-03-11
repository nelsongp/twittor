package models

/*RespuestaConsultaRelacion tiene el tru o flase que se obtiene al consultar la relacion entre 2 usuarios*/
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"`
}
