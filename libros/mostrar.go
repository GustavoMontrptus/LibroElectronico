package libros

import (
	"fmt"
)

func MostrarGenerosDisponibles() {
	fmt.Println("Géneros disponibles:")
	for _, genero := range genero {
		fmt.Printf("%d. %s\n", genero.ID, genero.Nombre)
	}
	fmt.Println("Ingrese 's' para salir del programa.")
}

func MostrarLibrosPorGenero(generoID int) bool {
	encontrado := false
	fmt.Println("Libros disponibles:")
	for _, libro := range libros {
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

//Darwin tandazo
//gggg
