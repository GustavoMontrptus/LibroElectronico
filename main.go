package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"Proyecto/autor"
	"Proyecto/editorial"
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
	log.Println("Entrando en indexHandler")
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Error cargando la plantilla index.html:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error ejecutando la plantilla index.html:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *App) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println("Entrando en loginHandler (GET)")
		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			log.Println("Error cargando la plantilla login.html:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println("Error ejecutando la plantilla login.html:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		log.Println("Entrando en loginHandler (POST)")
		r.ParseForm()
		correo := r.FormValue("correo")
		contrasena := r.FormValue("contrasena")

		_, autenticado := usuario.AutenticarUsuario(app.DB, correo, contrasena)
		if autenticado {
			http.Redirect(w, r, "/libros", http.StatusSeeOther)
		} else {
			err := usuario.CrearCliente(app.DB, correo, contrasena)
			if err != nil {
				log.Println("Error creando el usuario:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/libros", http.StatusSeeOther)
		}
	}
}

func (app *App) librosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println("Entrando en librosHandler (GET)")
		tmpl, err := template.ParseFiles("templates/seleccion_genero.html")
		if err != nil {
			log.Println("Error cargando la plantilla seleccion_genero.html:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		generos, err := genero.GetGeneros(app.DB)
		if err != nil {
			log.Println("Error obteniendo géneros:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Generos []genero.Genero
		}{
			Generos: generos,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println("Error ejecutando la plantilla seleccion_genero.html:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		log.Println("Entrando en librosHandler (POST)")
		generoID, err := strconv.Atoi(r.FormValue("genero"))
		if err != nil {
			log.Println("Error parsing genre ID:", err)
			http.Error(w, "Error parsing genre ID", http.StatusBadRequest)
			return
		}

		autores, err := autor.GetAutores(app.DB)
		if err != nil {
			log.Println("Error obteniendo autores:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		editoriales, err := editorial.GetEditoriales(app.DB)
		if err != nil {
			log.Println("Error obteniendo editoriales:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		generos, err := genero.GetGeneros(app.DB)
		if err != nil {
			log.Println("Error obteniendo géneros:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		success := libro.MostrarLibrosPorGenero(app.DB, generoID, autores, editoriales, generos)
		if !success {
			http.Error(w, "No se encontraron libros para el género seleccionado.", http.StatusNotFound)
			return
		}
	}
}

func (app *App) detalleLibroHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del libro desde los parámetros de la URL
	libroID := r.URL.Query().Get("id")

	// Consultar la base de datos para obtener los detalles del libro
	libroDetalles, err := libro.ObtenerDetallesLibro(app.DB, libroID)
	if err != nil {
		log.Println("Error obteniendo detalles del libro:", err)
		http.Error(w, "Error obteniendo detalles del libro", http.StatusInternalServerError)
		return
	}

	// Renderizar la plantilla de detalles del libro con los datos obtenidos
	tmpl, err := template.ParseFiles("templates/detalle_libro.html")
	if err != nil {
		log.Println("Error cargando la plantilla detalle_libro.html:", err)
		http.Error(w, "Error cargando la plantilla de detalles del libro", http.StatusInternalServerError)
		return
	}

	// Pasar los datos del libro a la plantilla
	err = tmpl.Execute(w, libroDetalles)
	if err != nil {
		log.Println("Error ejecutando la plantilla detalle_libro.html:", err)
		http.Error(w, "Error mostrando detalles del libro", http.StatusInternalServerError)
		return
	}
}

func main() {
	db, err := sql.Open("mysql", "root:camila262004@tcp(127.0.0.1:3306)/biblioteca_go")
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
	http.HandleFunc("/detalle-libro", app.detalleLibroHandler)

	// Rutas para los endpoints definidos en services.go
	serviceApp := &services.App{DB: db}
	http.HandleFunc("/autores", serviceApp.GetAutoresHandler)
	http.HandleFunc("/editoriales", serviceApp.GetEditorialesHandler)
	http.HandleFunc("/generos", serviceApp.GetGenerosHandler)
	http.HandleFunc("/servicio-libros", serviceApp.GetLibrosHandler)

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
