package genero

import "fmt"

type Genero struct {
	ID     int
	Nombre string
}
// Constructor
func NewGenero(ID int, Nombre string) *Genero {
	return &Genero{
		ID:     ID,
		Nombre: Nombre,
	}
}
// Getters
func (g *Genero) GetID() int {
	return g.ID
}

func (g *Genero) GetNombre() string {
	return g.Nombre
}
// Setters
func (g *Genero) SetID(ID int) {
	g.ID = ID
}

func (g *Genero) SetNombre(Nombre string) {
	g.Nombre = Nombre
}

// Método
func (g *Genero) PrintGenero() {
	fmt.Printf("ID: %d\nNombre: %s\n", g.ID, g.Nombre)
}
