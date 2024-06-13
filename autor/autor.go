package autor

import "fmt"

type Autor struct {
	ID           int
	Nombre       string
	Nacionalidad string
}

// Constructor
func NewAutor(ID int, Nombre string, Nacionalidad string) *Autor {
	return &Autor{
		ID:           ID,
		Nombre:       Nombre,
		Nacionalidad: Nacionalidad,
	}
}

// Getters
func (a *Autor) GetID() int {
	return a.ID
}

func (a *Autor) GetNombre() string {
	return a.Nombre
}

func (a *Autor) GetNacionalidad() string {
	return a.Nacionalidad
}

// Setters
func (a *Autor) SetID(ID int) {
	a.ID = ID
}

func (a *Autor) SetNombre(Nombre string) {
	a.Nombre = Nombre
}

func (a *Autor) SetNacionalidad(Nacionalidad string) {
	a.Nacionalidad = Nacionalidad
}

// Método
func (a *Autor) PrintAutor() {
	fmt.Printf("ID: %d\nNombre: %s\nNacionalidad: %s\n", a.ID, a.Nombre, a.Nacionalidad)
}
