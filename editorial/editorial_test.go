package editorial

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetEditoriales(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Editorial 1").
		AddRow(2, "Editorial 2")

	mock.ExpectQuery("SELECT id, name FROM editorial").WillReturnRows(rows)

	editoriales, err := GetEditoriales(db)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if len(editoriales) != 2 {
		t.Errorf("expected 2 editoriales, got %d", len(editoriales))
	}
}

func TestGetNameEditorialByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT name FROM editorial WHERE id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Editorial 1"))

	name, err := GetNameEditorialByID(db, 1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if name != "Editorial 1" {
		t.Errorf("expected 'Editorial 1', got '%s'", name)
	}
}
