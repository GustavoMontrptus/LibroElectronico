package usuario

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAutenticarUsuario(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, correo, contrasena FROM usuarios WHERE correo = ? AND contrasena = ?").WithArgs("usuario@example.com", "contraseña").
		WillReturnRows(sqlmock.NewRows([]string{"id", "correo", "contrasena"}).AddRow(1, "usuario@example.com", "contraseña"))

	_, autenticado := AutenticarUsuario(db, "usuario@example.com", "contraseña")
	if !autenticado {
		t.Errorf("expected user to be authenticated")
	}
}

func TestCrearCliente(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO usuarios").WithArgs("usuario@example.com", "contraseña").WillReturnResult(sqlmock.NewResult(1, 1))

	err = CrearCliente(db, "usuario@example.com", "contraseña")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}
