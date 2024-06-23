package autor

import (
	"database/sql"
	"fmt"
)

type Autor struct {
	ID           int
	Nombre       string
	Nacionalidad string
}

func GetAutores(db *sql.DB) ([]Autor, error) {
	rows, err := db.Query("SELECT id, name, nacionalidad FROM autor")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var autores []Autor
	for rows.Next() {
		var autor Autor
		if err := rows.Scan(&autor.ID, &autor.Nombre, &autor.Nacionalidad); err != nil {
			return nil, err
		}
		autores = append(autores, autor)
	}
	return autores, nil
}

func (a *Autor) PrintAutor() {
	fmt.Printf("ID: %d\nNombre: %s\nNacionalidad: %s\n", a.ID, a.Nombre, a.Nacionalidad)
}
