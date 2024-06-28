package autor

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAutores(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "nacionalidad"}).
		AddRow(1, "Autor 1", "Nacionalidad 1").
		AddRow(2, "Autor 2", "Nacionalidad 2")

	mock.ExpectQuery("SELECT id, name, nacionalidad FROM autor").WillReturnRows(rows)

	autores, err := GetAutores(db)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if len(autores) != 2 {
		t.Errorf("expected 2 autores, got %d", len(autores))
	}
}

// funcion para buscar autor por id
func TestGetNameAutorByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT name FROM autor WHERE id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Autor 1"))

	name, err := GetNameAutorByID(db, 1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if name != "Autor 1" {
		t.Errorf("expected 'Autor 1', got '%s'", name)
	}
}
