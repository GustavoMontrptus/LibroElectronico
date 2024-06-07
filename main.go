package main

import (
	"bufio"
	"fmt"
	"miaplicacion/usuario"
	"os"
	"strings"
)

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
			usuario.CrearNuevaCuenta(reader)
		case "2":
			usuario.IniciarSesion(reader)
		case "3":
			fmt.Println("Saliendo del programa...")
			os.Exit(0)
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

// prueba git
//MIKKE

// Prueba Git
//Bryan Espinoza
