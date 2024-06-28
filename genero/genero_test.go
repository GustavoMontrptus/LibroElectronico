package genero

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetGeneros(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Genero 1").
		AddRow(2, "Genero 2")

	mock.ExpectQuery("SELECT id, name FROM genero").WillReturnRows(rows)

	generos, err := GetGeneros(db)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if len(generos) != 2 {
		t.Errorf("expected 2 generos, got %d", len(generos))
	}
}

func TestGetNameGeneroByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT name FROM genero WHERE id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Genero 1"))

	name, err := GetNameGeneroByID(db, 1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if name != "Genero 1" {
		t.Errorf("expected 'Genero 1', got '%s'", name)
	}
}
