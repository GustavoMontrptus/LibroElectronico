package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"Proyecto/genero"
	"Proyecto/libro"
	"Proyecto/services"
	"Proyecto/usuario"

	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	DB *sql.DB
}

func (app *App) indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func (app *App) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		correo := r.FormValue("correo")
		contrasena := r.FormValue("contrasena")
		accion := r.FormValue("accion")

		switch accion {
		case "login":
			if _, autenticado := usuario.AutenticarUsuario(app.DB, correo, contrasena); autenticado {
				http.Redirect(w, r, "/libros", http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/login?error=credenciales_invalidas", http.StatusSeeOther)
			}
		case "registro":
			if err := usuario.CrearCliente(app.DB, correo, contrasena); err != nil {
				log.Printf("Error al crear usuario: %v", err)
				http.Error(w, "Error al crear usuario: "+err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/libros", http.StatusSeeOther)
		default:
			http.Error(w, "Acción no válida", http.StatusBadRequest)
		}
	}
}

func (app *App) librosHandler(w http.ResponseWriter, r *http.Request) {
	generoID := r.URL.Query().Get("genero")

	if generoID == "" {
		generos, err := genero.GetGeneros(app.DB)
		if err != nil {
			http.Error(w, "Error al obtener géneros", http.StatusInternalServerError)
			return
		}

		tmpl := template.Must(template.ParseFiles("templates/seleccion_genero.html"))
		tmpl.Execute(w, generos)
		return
	}

	generoIDInt, err := strconv.Atoi(generoID)
	if err != nil {
		http.Error(w, "ID de género inválido", http.StatusBadRequest)
		return
	}

	libros, err := libro.ObtenerLibrosPorGenero(app.DB, generoIDInt)
	if err != nil {
		http.Error(w, "Error al obtener libros", http.StatusInternalServerError)
		log.Printf("Error obteniendo libros: %v", err)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/lista_libros.html"))
	tmpl.Execute(w, libros)
}

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/biblioteca")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	app := &App{
		DB: db,
	}

	http.HandleFunc("/", app.indexHandler)
	http.HandleFunc("/login", app.loginHandler)
	http.HandleFunc("/libros", app.librosHandler)

	serviceApp := &services.App{DB: db}
	http.HandleFunc("/autores", serviceApp.GetAutoresHandler)
	http.HandleFunc("/editoriales", serviceApp.GetEditorialesHandler)
	http.HandleFunc("/generos", serviceApp.GetGenerosHandler)
	http.HandleFunc("/servicio-libros", serviceApp.GetLibrosHandler)

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
