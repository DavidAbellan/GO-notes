package user

import "time"

type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

//this significa que haces referencia a esa estructura
func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool) {
	this.FechaAlta = fechaalta
	this.Id = id
	this.Nombre = nombre
	this.Status = status

}
