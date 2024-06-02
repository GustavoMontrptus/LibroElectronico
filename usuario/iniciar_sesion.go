package usuario

import (
	"bufio"
	"fmt"
	"miaplicacion/libros"
	"os"
	"strings"
)

func IniciarSesion(reader *bufio.Reader) {
	fmt.Print("Ingrese su correo: ")
	correo, _ := reader.ReadString('\n')
	correo = strings.TrimSpace(correo)

	fmt.Print("Ingrese la contraseña: ")
	contrasena, _ := reader.ReadString('\n')
	contrasena = strings.TrimSpace(contrasena)

	usuario, autenticado := AutenticarUsuario(correo, contrasena)
	if autenticado {
		fmt.Printf("¡Bienvenido, %s!\n", usuario.Correo)
		for {
			libros.MostrarGenerosDisponibles()
			fmt.Print("Seleccione un género por su número o 's' para salir: ")
			generoIDStr, _ := reader.ReadString('\n')
			generoIDStr = strings.TrimSpace(generoIDStr)
			if generoIDStr == "s" {
				fmt.Println("Saliendo del programa...")
				os.Exit(0)
			}

			var generoID int
			fmt.Sscanf(generoIDStr, "%d", &generoID)

			if libros.MostrarLibrosPorGenero(generoID) {
				fmt.Print("Ingrese 'r' para regresar a la lista de géneros o 's' para salir: ")
				opcion, _ := reader.ReadString('\n')
				opcion = strings.TrimSpace(opcion)
				if opcion == "r" {
					continue
				} else if opcion == "s" {
					fmt.Println("Saliendo del programa...")
					os.Exit(0)
				}
				continue
			}
		}
	} else {
		fmt.Println("Correo o contraseña incorrecta.")
	}
}
