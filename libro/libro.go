package libro

import (
	"Proyecto/autor"
	"Proyecto/editorial"
	"Proyecto/genero"
	"fmt"
)

type Libro struct {
	Titulo      string
	GeneroID    int
	AutorID     int
	Year        int
	Descripcion string
	EditorialID int
}

var Libros = []Libro{
	{"El señor de los anillos", 1, 1, 1954, "Un libro de fantasía épica.", 1},
	{"Cien años de soledad", 2, 2, 1967, "Una obra maestra del realismo mágico.", 2},
	{"Don Quijote de la Mancha", 3, 3, 1605, "Una novela clásica de la literatura española.", 3},
	{"Yo Robot", 4, 4, 1950, "Es una de las obras más populares de Asimov.", 4},
	{"El resplandor", 5, 5, 1977, "La maldad acecha las mentes débiles.", 5},
	{"Lisístrata", 6, 6, 2003, "Busco la comedia en los momentos más oscuros.", 6},
	{"Cometas en el cielo", 7, 7, 2003, "Narra el ir y venir en la vida de dos niños.", 7},
}

// Getters
func (l *Libro) GetTitulo() string {
	return l.Titulo
}

func (l *Libro) GetGeneroID() int {
	return l.GeneroID
}

func (l *Libro) GetAutorID() int {
	return l.AutorID
}

func (l *Libro) GetYear() int {
	return l.Year
}

func (l *Libro) GetDescripcion() string {
	return l.Descripcion
}

func (l *Libro) GetEditorialID() int {
	return l.EditorialID
}

// Setters
func (l *Libro) SetTitulo(Titulo string) {
	l.Titulo = Titulo
}

func (l *Libro) SetGeneroID(GeneroID int) {
	l.GeneroID = GeneroID
}

func (l *Libro) SetAutorID(AutorID int) {
	l.AutorID = AutorID
}

func (l *Libro) SetYear(Year int) {
	l.Year = Year
}

func (l *Libro) SetDescripcion(Descripcion string) {
	l.Descripcion = Descripcion
}

func (l *Libro) SetEditorialID(EditorialID int) {
	l.EditorialID = EditorialID
}

func MostrarLibrosPorGenero(generoID int, autores []autor.Autor, editoriales []editorial.Editorial, generos []genero.Genero) bool {
	encontrado := false
	fmt.Println("Libros disponibles:")
	for _, libro := range Libros {
		if libro.GeneroID == generoID {
			encontrado = true
			autor := autores[libro.AutorID-1]
			editorial := editoriales[libro.EditorialID-1]
			genero := generos[libro.GeneroID-1]
			fmt.Printf("Título: %s\nAño: %d\nDescripción: %s\nGénero: %s\nAutor: %s\nNacionalidad del Autor: %s\nEditorial: %s\n----\n",
				libro.Titulo, libro.Year, libro.Descripcion, genero.Nombre, autor.Nombre, autor.Nacionalidad, editorial.Nombre)
		}
	}
	if !encontrado {
		fmt.Println("No hay libros disponibles en este género.")
	}
	if encontrado {
		fmt.Println("Ingrese 'r' para regresar a la lista de géneros o 's' para salir del programa.")
	}
	return encontrado
}
