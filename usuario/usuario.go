package usuario

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Usuario struct {
	correo     string
	contrasena string
}

var UsuariosRegistrados = make(map[string]*Usuario)

// Getters
func (u *Usuario) GetCorreo() string {
	return u.correo
}

func (u *Usuario) GetContrasena() string {
	return u.contrasena
}

// Setters
func (u *Usuario) SetCorreo(correo string) {
	u.correo = correo
}

func (u *Usuario) SetContrasena(contrasena string) {
	u.contrasena = contrasena
}

func GuardarUsuarios() {
	file, err := os.Create("usuarios.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	for _, usuario := range UsuariosRegistrados {
		fmt.Fprintf(file, "%s,%s\n", usuario.GetCorreo(), usuario.GetContrasena())
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
			usuario := &Usuario{}
			usuario.SetCorreo(partes[0])
			usuario.SetContrasena(partes[1])
			UsuariosRegistrados[partes[0]] = usuario
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}
}

func CrearCliente(correo, contrasena string) {
	usuario := &Usuario{}
	usuario.SetCorreo(correo)
	usuario.SetContrasena(contrasena)
	UsuariosRegistrados[correo] = usuario
	GuardarUsuarios()
}

func AutenticarUsuario(correo, contrasena string) (*Usuario, bool) {
	usuario, existe := UsuariosRegistrados[correo]
	if !existe || usuario.GetContrasena() != contrasena {
		return nil, false
	}
	return usuario, true
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
