package usuario

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Usuario struct {
	Correo     string
	Contrasena string
}

var usuariosRegistrados = make(map[string]*Usuario)

func GuardarUsuarios() {
	file, err := os.Create("usuarios.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	for _, usuario := range usuariosRegistrados {
		fmt.Fprintf(file, "%s,%s\n", usuario.Correo, usuario.Contrasena)
	}
}

func CargarUsuarios() {
	file, err := os.Open("usuarios.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		partes := strings.Split(linea, ",")
		if len(partes) == 2 {
			usuariosRegistrados[partes[0]] = &Usuario{Correo: partes[0], Contrasena: partes[1]}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}
}

func CrearCliente(correo, contrasena, nombreTitular, numeroTarjeta, fechaCaducidad, claveSeguridad string) {
	usuario := &Usuario{
		Correo:     correo,
		Contrasena: contrasena,
	}
	usuariosRegistrados[correo] = usuario
	GuardarUsuarios()
}

func AutenticarUsuario(correo, contrasena string) (*Usuario, bool) {
	usuario, existe := usuariosRegistrados[correo]
	if !existe || usuario.Contrasena != contrasena {
		return nil, false
	}
	return usuario, true
}

func validarCorreo(correo string) bool {
	if strings.Contains(correo, "@") && strings.HasSuffix(correo, ".com") {
		return true
	}
	return false
}

func validarContrasena(contrasena string) bool {
	if len(contrasena) >= 8 {
		for _, char := range contrasena {
			if unicode.IsUpper(char) {
				return true
			}
		}
	}
	return false
}
