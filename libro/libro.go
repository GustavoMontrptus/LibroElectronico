package libro

import (
	"Proyecto/autor"
	"Proyecto/editorial"
	"Proyecto/genero"
	"database/sql"
	"fmt"
)

type Libro struct {
	ID              int
	Titulo          string
	GeneroID        int
	AutorID         int
	EditorialID     int
	Year            int
	Descripcion     string
	NombreGenero    string
	NombreAutor     string
	NombreEditorial string
}

func GetLibros(db *sql.DB) ([]Libro, error) {
	query := "SELECT id, titulo, genero_id, autor_id, editorial_id, year, descripcion FROM libro"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var libros []Libro
	for rows.Next() {
		var l Libro

		genero, err := genero.GetNameGeneroByID(db, l.GeneroID)
		if err != nil {
			return nil, err
		}

		if err := rows.Scan(&l.ID, &l.Titulo, &genero, &l.AutorID, &l.EditorialID, &l.Year, &l.Descripcion); err != nil {
			return nil, err
		}

		libros = append(libros, l)
	}
	return libros, nil
}

func ObtenerLibrosPorGenero(db *sql.DB, generoID int) ([]Libro, error) {
	query := "SELECT id, titulo, genero_id, autor_id, editorial_id, year, descripcion FROM libro WHERE genero_id = ?"
	rows, err := db.Query(query, generoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	genero, errGenero := genero.GetNameGeneroByID(db, generoID)
	if errGenero != nil {
		return nil, errGenero
	}

	var libros []Libro
	for rows.Next() {
		var l Libro
		if err := rows.Scan(&l.ID, &l.Titulo, &l.GeneroID, &l.AutorID, &l.EditorialID, &l.Year, &l.Descripcion); err != nil {
			return nil, err
		}

		autor, errAutor := autor.GetNameAutorByID(db, l.AutorID)
		if errAutor != nil {
			return nil, errAutor
		}

		editorial, errEditorial := editorial.GetNameEditorialByID(db, l.EditorialID)
		if errEditorial != nil {
			return nil, errEditorial
		}

		l.NombreGenero = genero
		l.NombreAutor = autor
		l.NombreEditorial = editorial

		libros = append(libros, l)
	}

	// Verificar errores después del loop
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return libros, nil
}

func MostrarLibrosPorGenero(db *sql.DB, generoID int, autores []autor.Autor, editoriales []editorial.Editorial, generos []genero.Genero) bool {
	libros, err := ObtenerLibrosPorGenero(db, generoID)
	if err != nil {
		fmt.Println("Error obteniendo libros:", err)
		return false
	}

	if len(libros) == 0 {
		fmt.Println("No se encontraron libros para el género seleccionado.")
		return true
	}

	fmt.Println("Libros disponibles:")
	for _, l := range libros {
		fmt.Printf("ID: %d\n", l.ID)
		fmt.Printf("Título: %s\n", l.Titulo)
		fmt.Printf("Año: %d\n", l.Year)
		fmt.Printf("Descripción: %s\n", l.Descripcion)

		for _, a := range autores {
			if a.ID == l.AutorID {
				fmt.Printf("Autor: %s\n", a.Nombre)
				break
			}
		}

		for _, e := range editoriales {
			if e.ID == l.EditorialID {
				fmt.Printf("Editorial: %s\n", e.Nombre)
				break
			}
		}

		for _, g := range generos {
			if g.ID == l.GeneroID {
				fmt.Printf("Género: %s\n", g.Nombre)
				break
			}
		}
		fmt.Println("----------------------------")
	}

	return true
}
