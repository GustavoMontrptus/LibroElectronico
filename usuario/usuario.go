package usuario

import (
	"database/sql"
	"fmt"
	"strings"
	"unicode"
)

type Usuario struct {
	ID         int
	Correo     string
	Contrasena string
}

func CrearCliente(db *sql.DB, correo, contrasena string) error {
	_, err := db.Exec("INSERT INTO usuarios (correo, contrasena) VALUES (?, ?)", correo, contrasena)
	return err
}

func AutenticarUsuario(db *sql.DB, correo, contrasena string) (*Usuario, bool) {
	var usuario Usuario
	err := db.QueryRow("SELECT id, correo, contrasena FROM usuarios WHERE correo = ? AND contrasena = ?", correo, contrasena).Scan(&usuario.ID, &usuario.Correo, &usuario.Contrasena)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false
		}
		fmt.Println("Error querying usuario:", err)
		return nil, false
	}
	return &usuario, true
}

func ValidarCorreo(correo string) bool {
	return strings.Contains(correo, "@") && strings.HasSuffix(correo, ".com")
}

func ValidarContrasena(contrasena string) bool {
	if len(contrasena) < 8 {
		return false
	}
	for _, char := range contrasena {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}
