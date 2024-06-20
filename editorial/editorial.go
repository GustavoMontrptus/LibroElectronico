package editorial

import "fmt"

type Editorial struct {
	ID     int
	Nombre string
}

// Constructor
func NewEditorial(ID int, Nombre string) *Editorial {
	return &Editorial{
		ID:     ID,
		Nombre: Nombre,
	}
}
// Getters
func (e *Editorial) GetID() int {
	return e.ID
}

func (e *Editorial) GetNombre() string {
	return e.Nombre
}

// Setters
func (e *Editorial) SetID(ID int) {
	e.ID = ID
}

func (e *Editorial) SetNombre(Nombre string) {
	e.Nombre = Nombre
}
// Método
func (e *Editorial) PrintEditorial() {
	fmt.Printf("ID: %d\nNombre: %s\n", e.ID, e.Nombre)
}
