package main

import (
	"Proyecto/autor"
	"Proyecto/editorial"
	"Proyecto/genero"
	"Proyecto/libro"
	"Proyecto/usuario"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:camila262004@tcp(127.0.0.1:3306)/biblioteca_go")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database successfully!")
}

func main() {
	initDB()
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese su correo:")
	scanner.Scan()
	correo := scanner.Text()

	fmt.Println("Ingrese su contraseña:")
	scanner.Scan()
	contrasena := scanner.Text()

	if !usuario.ValidarCorreo(correo) {
		fmt.Println("Correo inválido. Debe contener '@' y terminar en '.com'")
		return
	}

	if !usuario.ValidarContrasena(contrasena) {
		fmt.Println("Contraseña inválida. Debe tener al menos 8 caracteres y una mayúscula.")
		return
	}

	user, autenticado := usuario.AutenticarUsuario(db, correo, contrasena)
	if autenticado {
		fmt.Printf("Bienvenido, %s!\n", user.Correo)
	} else {
		fmt.Println("Usuario no encontrado. Creando nuevo usuario...")
		err := usuario.CrearCliente(db, correo, contrasena)
		if err != nil {
			fmt.Println("Error creando usuario:", err)
			return
		}
		fmt.Println("Usuario creado exitosamente!")
	}

	autores, err := autor.GetAutores(db)
	if err != nil {
		fmt.Println("Error obteniendo autores:", err)
		return
	}

	editoriales, err := editorial.GetEditoriales(db)
	if err != nil {
		fmt.Println("Error obteniendo editoriales:", err)
		return
	}

	generos, err := genero.GetGeneros(db)
	if err != nil {
		fmt.Println("Error obteniendo generos:", err)
		return
	}

	mostrarGeneros(generos)

	for {
		fmt.Println("Seleccione el ID del género que desea ver:")
		scanner.Scan()
		generoID, _ := strconv.Atoi(scanner.Text())

		if libro.MostrarLibrosPorGenero(db, generoID, autores, editoriales, generos) {
			fmt.Println("Ingrese 'r' para regresar a la lista de géneros o 's' para salir del programa.")
			scanner.Scan()
			decision := scanner.Text()
			if decision == "s" {
				break
			}
			if decision == "r" {
				mostrarGeneros(generos)
			}
		}
	}
}

func mostrarGeneros(generos []genero.Genero) {
	fmt.Println("Géneros disponibles:")
	for _, g := range generos {
		fmt.Printf("%d. %s\n", g.ID, g.Nombre)
	}
}
