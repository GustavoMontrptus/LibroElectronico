package libro

import (
	"Proyecto/autor"
	"Proyecto/editorial"
	"Proyecto/genero"
	"database/sql"
	"fmt"
)

type Libro struct {
	ID          int
	Titulo      string
	GeneroID    int
	AutorID     int
	EditorialID int
	Year        int
	Descripcion string
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
		if err := rows.Scan(&l.ID, &l.Titulo, &l.GeneroID, &l.AutorID, &l.EditorialID, &l.Year, &l.Descripcion); err != nil {
			return nil, err
		}
		libros = append(libros, l)
	}
	return libros, nil
}

func MostrarLibrosPorGenero(db *sql.DB, generoID int, autores []autor.Autor, editoriales []editorial.Editorial, generos []genero.Genero) bool {
	query := "SELECT id, titulo, autor_id, editorial_id, year, descripcion FROM libro WHERE genero_id = ?"
	rows, err := db.Query(query, generoID)
	if err != nil {
		fmt.Println("Error querying libros:", err)
		return false
	}
	defer rows.Close()

	fmt.Println("Libros disponibles:")
	found := false
	for rows.Next() {
		found = true
		var l Libro
		if err := rows.Scan(&l.ID, &l.Titulo, &l.AutorID, &l.EditorialID, &l.Year, &l.Descripcion); err != nil {
			fmt.Println("Error scanning libro:", err)
			return false
		}
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

	if !found {
		fmt.Println("No se encontraron libros para el género seleccionado.")
	}

	return true
}

func ObtenerDetallesLibro(db *sql.DB, libroID string) (*Libro, error) {
	query := "SELECT id, titulo, genero_id, autor_id, editorial_id, year, descripcion FROM libro WHERE id = ?"
	row := db.QueryRow(query, libroID)

	var l Libro
	err := row.Scan(&l.ID, &l.Titulo, &l.GeneroID, &l.AutorID, &l.EditorialID, &l.Year, &l.Descripcion)
	if err != nil {
		return nil, err
	}

	return &l, nil
}
