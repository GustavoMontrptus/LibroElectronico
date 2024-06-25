package usuario

import (
	"database/sql"
	"fmt"
	"regexp"
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
	err := db.QueryRow("SELECT id, correo, contrasena FROM usuarios WHERE correo = ? AND contrasena = ?", correo, contrasena).
		Scan(&usuario.ID, &usuario.Correo, &usuario.Contrasena)

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
	// Validación básica del formato de correo
	// Debe contener al menos un arroba y un punto seguido de "com"
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(correo)
}

func ValidarContrasena(contrasena string) bool {
	// Validación de la contraseña
	// Debe tener al menos 8 caracteres, una mayúscula y un número
	if len(contrasena) < 8 {
		return false
	}
	hasUpperCase := false
	hasDigit := false
	for _, char := range contrasena {
		if 'A' <= char && char <= 'Z' {
			hasUpperCase = true
		}
		if '0' <= char && char <= '9' {
			hasDigit = true
		}
	}
	return hasUpperCase && hasDigit
}
