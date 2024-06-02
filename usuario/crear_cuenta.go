package usuario

import (
	"bufio"
	"fmt"
	"strings"
)

func CrearNuevaCuenta(reader *bufio.Reader) {
	var correo, contrasena string
	for {
		fmt.Print("Ingrese su correo: ")
		correo, _ = reader.ReadString('\n')
		correo = strings.TrimSpace(correo)
		if validarCorreo(correo) {
			break
		} else {
			fmt.Println("Correo inválido. Debe contener un arroba y terminar con '.com'.")
		}
	}

	for {
		fmt.Print("Ingrese la contraseña: ")
		contrasena, _ = reader.ReadString('\n')
		contrasena = strings.TrimSpace(contrasena)
		if validarContrasena(contrasena) {
			break
		} else {
			fmt.Println("Contraseña inválida. Debe tener al menos 8 caracteres y contener al menos una letra mayúscula.")
		}
	}

	fmt.Print("Ingrese el nombre del titular de la tarjeta: ")
	nombreTitular, _ := reader.ReadString('\n')
	nombreTitular = strings.TrimSpace(nombreTitular)

	fmt.Print("Ingrese el número de tarjeta: ")
	numeroTarjeta, _ := reader.ReadString('\n')
	numeroTarjeta = strings.TrimSpace(numeroTarjeta)

	fmt.Print("Ingrese la fecha de caducidad: ")
	fechaCaducidad, _ := reader.ReadString('\n')
	fechaCaducidad = strings.TrimSpace(fechaCaducidad)

	fmt.Print("Ingrese la clave de seguridad: ")
	claveSeguridad, _ := reader.ReadString('\n')
	claveSeguridad = strings.TrimSpace(claveSeguridad)

	CrearCliente(correo, contrasena, nombreTitular, numeroTarjeta, fechaCaducidad, claveSeguridad)
	fmt.Println("Cliente creado exitosamente!")
}
