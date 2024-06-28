package libro

import (
	"Proyecto/autor"
	"Proyecto/editorial"
	"Proyecto/genero"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetLibros(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Definir filas simuladas para la consulta SELECT
	rows := sqlmock.NewRows([]string{"id", "titulo", "genero_id", "autor_id", "editorial_id", "year", "descripcion"}).
		AddRow(1, "Libro 1", 1, 1, 1, 2020, "Descripción 1").
		AddRow(2, "Libro 2", 2, 2, 2, 2019, "Descripción 2")

	// Expectativa de consulta SQL
	mock.ExpectQuery("SELECT id, titulo, genero_id, autor_id, editorial_id, year, descripcion FROM libro").WillReturnRows(rows)

	// Expectativas adicionales para consultas de género, autor y editorial si existen
	mock.ExpectQuery("SELECT name FROM genero WHERE id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Genero 1"))
	mock.ExpectQuery("SELECT name FROM genero WHERE id = ?").WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Genero 2"))
	mock.ExpectQuery("SELECT name FROM autor WHERE id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Autor 1"))
	mock.ExpectQuery("SELECT name FROM autor WHERE id = ?").WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Autor 2"))
	mock.ExpectQuery("SELECT name FROM editorial WHERE id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Editorial 1"))
	mock.ExpectQuery("SELECT name FROM editorial WHERE id = ?").WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Editorial 2"))

	// Ejecutar la función bajo prueba
	libros, err := GetLibros(db)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	// Verificar resultados
	if len(libros) != 2 {
		t.Errorf("expected 2 libros, got %d", len(libros))
	}
}

func TestObtenerLibrosPorGenero(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	generoID := 1

	// Definir filas simuladas para la consulta SELECT con filtro por género
	rows := sqlmock.NewRows([]string{"id", "titulo", "genero_id", "autor_id", "editorial_id", "year", "descripcion"}).
		AddRow(1, "Libro 1", generoID, 1, 1, 2020, "Descripción 1").
		AddRow(2, "Libro 2", generoID, 2, 2, 2019, "Descripción 2")

	// Expectativa de consulta SQL
	mock.ExpectQuery("SELECT id, titulo, genero_id, autor_id, editorial_id, year, descripcion FROM libro WHERE genero_id = ?").WithArgs(generoID).WillReturnRows(rows)

	// Expectativas adicionales para consultas de género, autor y editorial si existen
	mock.ExpectQuery("SELECT name FROM autor WHERE id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Autor 1"))
	mock.ExpectQuery("SELECT name FROM autor WHERE id = ?").WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Autor 2"))
	mock.ExpectQuery("SELECT name FROM editorial WHERE id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Editorial 1"))
	mock.ExpectQuery("SELECT name FROM editorial WHERE id = ?").WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Editorial 2"))

	// Ejecutar la función bajo prueba
	libros, err := ObtenerLibrosPorGenero(db, generoID)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	// Verificar resultados
	if len(libros) != 2 {
		t.Errorf("expected 2 libros, got %d", len(libros))
	}
}

func TestMostrarLibrosPorGenero(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	generoID := 1

	// Definir filas simuladas para la consulta SELECT con filtro por género
	rows := sqlmock.NewRows([]string{"id", "titulo", "genero_id", "autor_id", "editorial_id", "year", "descripcion"}).
		AddRow(1, "Libro 1", generoID, 1, 1, 2020, "Descripción 1").
		AddRow(2, "Libro 2", generoID, 2, 2, 2019, "Descripción 2")

	// Expectativa de consulta SQL
	mock.ExpectQuery("SELECT id, titulo, genero_id, autor_id, editorial_id, year, descripcion FROM libro WHERE genero_id = ?").WithArgs(generoID).WillReturnRows(rows)

	// Mocks adicionales para autores, editoriales y géneros
	autores := []autor.Autor{{ID: 1, Nombre: "Autor 1"}, {ID: 2, Nombre: "Autor 2"}}
	editoriales := []editorial.Editorial{{ID: 1, Nombre: "Editorial 1"}, {ID: 2, Nombre: "Editorial 2"}}
	generos := []genero.Genero{{ID: 1, Nombre: "Genero 1"}, {ID: 2, Nombre: "Genero 2"}}

	// Ejecutar la función bajo prueba
	ok := MostrarLibrosPorGenero(db, generoID, autores, editoriales, generos)
	if !ok {
		t.Errorf("expected MostrarLibrosPorGenero to return true")
	}
}
