package genero

import (
	"database/sql"
	"fmt"
)

type Genero struct {
	ID     int
	Nombre string
}

func GetGeneros(db *sql.DB) ([]Genero, error) {
	rows, err := db.Query("SELECT id, name FROM genero")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var generos []Genero
	for rows.Next() {
		var genero Genero
		if err := rows.Scan(&genero.ID, &genero.Nombre); err != nil {
			return nil, err
		}
		generos = append(generos, genero)
	}
	return generos, nil
}

func (g *Genero) PrintGenero() {
	fmt.Printf("ID: %d\nNombre: %s\n", g.ID, g.Nombre)
}
