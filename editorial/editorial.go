package editorial

import (
	"database/sql"
	"fmt"
)

type Editorial struct {
	ID     int
	Nombre string
}

func GetEditoriales(db *sql.DB) ([]Editorial, error) {
	rows, err := db.Query("SELECT id, name FROM editorial")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var editoriales []Editorial
	for rows.Next() {
		var editorial Editorial
		if err := rows.Scan(&editorial.ID, &editorial.Nombre); err != nil {
			return nil, err
		}
		editoriales = append(editoriales, editorial)
	}
	return editoriales, nil
}

// funcion para buscar editorial por id
func GetNameEditorialByID(db *sql.DB, id int) (string, error) {
	var editorial string

	query := "SELECT name FROM editorial WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&editorial)
	return editorial, err
}

func (e *Editorial) PrintEditorial() {
	fmt.Printf("ID: %d\nNombre: %s\n", e.ID, e.Nombre)
}
