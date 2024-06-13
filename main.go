package main

import (
	"Proyecto/autor"
	"Proyecto/editorial"
	"Proyecto/genero"
	"Proyecto/libro"
	"Proyecto/usuario"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var generos = []genero.Genero{
	{ID: 1, Nombre: "Fantasía"},
	{ID: 2, Nombre: "Realismo Mágico"},
	{ID: 3, Nombre: "Clásicos"},
	{ID: 4, Nombre: "Ciencia Ficción"},
	{ID: 5, Nombre: "Terror"},
	{ID: 6, Nombre: "Comedia"},
	{ID: 7, Nombre: "Drama"},
}

var autores = []autor.Autor{
	{ID: 1, Nombre: "J.R.R. Tolkien", Nacionalidad: "Inglés"},
	{ID: 2, Nombre: "Gabriel García Márquez", Nacionalidad: "Colombiano"},
	{ID: 3, Nombre: "Miguel de Cervantes", Nacionalidad: "Español"},
	{ID: 4, Nombre: "Isaac Asimov", Nacionalidad: "Estadounidense"},
	{ID: 5, Nombre: "Stephen King", Nacionalidad: "Estadounidense"},
	{ID: 6, Nombre: "Aristófanes", Nacionalidad: "Griego"},
	{ID: 7, Nombre: "Khaled Hosseini", Nacionalidad: "Persa"},
}

var editoriales = []editorial.Editorial{
	{ID: 1, Nombre: "Allen & Unwin"},
	{ID: 2, Nombre: "Editorial Sudamericana"},
	{ID: 3, Nombre: "Francisco de Robles"},
	{ID: 4, Nombre: "Alianza Editorial"},
	{ID: 5, Nombre: "On Writing"},
	{ID: 6, Nombre: "Losada"},
	{ID: 7, Nombre: "Riverhead Books"},
}

func MostrarGenerosDisponibles() {
	fmt.Println("Géneros disponibles:")
	for _, genero := range generos {
		fmt.Printf("%d. %s\n", genero.ID, genero.Nombre)
	}
}

func main() {
	usuario.CargarUsuarios()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("¿Qué desea hacer?")
		fmt.Println("1. Crear una nueva cuenta de cliente")
		fmt.Println("2. Iniciar sesión")
		fmt.Println("3. Salir")
		fmt.Print("Ingrese su opción: ")
		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "1":
			var correo, contrasena string
			for {
				fmt.Print("Ingrese su correo: ")
				correo, _ = reader.ReadString('\n')
				correo = strings.TrimSpace(correo)
				if usuario.ValidarCorreo(correo) {
					break
				} else {
					fmt.Println("Correo inválido. Debe contener un arroba y terminar con '.com'.")
				}
			}

			for {
				fmt.Print("Ingrese la contraseña: ")
				contrasena, _ = reader.ReadString('\n')
				contrasena = strings.TrimSpace(contrasena)
				if usuario.ValidarContrasena(contrasena) {
					break
				} else {
					fmt.Println("Contraseña inválida. Debe tener al menos 8 caracteres y contener al menos una letra mayúscula.")
				}
			}

			usuario.CrearCliente(correo, contrasena)
			fmt.Println("Cliente creado exitosamente!")

		case "2":
			fmt.Print("Ingrese su correo: ")
			correo, _ := reader.ReadString('\n')
			correo = strings.TrimSpace(correo)

			fmt.Print("Ingrese la contraseña: ")
			contrasena, _ := reader.ReadString('\n')
			contrasena = strings.TrimSpace(contrasena)

			usr, autenticado := usuario.AutenticarUsuario(correo, contrasena)
			if autenticado {
				fmt.Printf("¡Bienvenido, %s!\n", usr.GetCorreo())
				for {
					MostrarGenerosDisponibles()
					fmt.Print("Seleccione un género por su número o 's' para salir: ")
					generoIDStr, _ := reader.ReadString('\n')
					generoIDStr = strings.TrimSpace(generoIDStr)
					if generoIDStr == "s" {
						fmt.Println("Saliendo del programa...")
						os.Exit(0)
					}

					var generoID int
					fmt.Sscanf(generoIDStr, "%d", &generoID)

					if libro.MostrarLibrosPorGenero(generoID, autores, editoriales, generos) {
						for {
							fmt.Print("Ingrese 'r' para regresar a la lista de géneros o 's' para salir: ")
							opcion, _ := reader.ReadString('\n')
							opcion = strings.TrimSpace(opcion)
							if opcion == "r" {
								break
							} else if opcion == "s" {
								fmt.Println("Saliendo del programa...")
								os.Exit(0)
							}
						}
					} else {
						fmt.Println("Opción no válida.")
					}
				}
			} else {
				fmt.Println("Correo o contraseña incorrecta.")
			}

		case "3":
			fmt.Println("Saliendo del programa...")
			os.Exit(0)

		default:
			fmt.Println("Opción no válida.")
		}
	}
}
