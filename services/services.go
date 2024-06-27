package services

import (
	"Proyecto/genero"
	"Proyecto/libro"
	"Proyecto/usuario"
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

type App struct {
	DB *sql.DB
}

func (app *App) IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func (app *App) LoginHandler(w http.ResponseWriter, r *http.Request) {
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
			_, autenticado := usuario.AutenticarUsuario(app.DB, correo, contrasena)
			if autenticado {
				http.Redirect(w, r, "/libros", http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/login?error=credencialesinvalidas", http.StatusSeeOther)
			}
		case "registro":
			if err := usuario.CrearCliente(app.DB, correo, contrasena); err != nil {
				http.Error(w, "Error al crear usuario", http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/libros", http.StatusSeeOther)
		default:
			http.Error(w, "Acción no válida", http.StatusBadRequest)
		}
	}
}

func (app *App) LibrosHandler(w http.ResponseWriter, r *http.Request) {
	generoID := r.URL.Query().Get("genero")
	if generoID == "" {
		generos, err := genero.GetGeneros(app.DB)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl := template.Must(template.ParseFiles("templates/seleccion_genero.html"))
		tmpl.Execute(w, generos)
		return
	}
	id, err := strconv.Atoi(generoID)
	if err != nil {
		http.Error(w, "ID de género inválido", http.StatusBadRequest)
		return
	}
	libros, err := libro.ObtenerLibrosPorGenero(app.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/libro.html"))
	tmpl.Execute(w, libros)
}

func (app *App) GetGenerosHandler(w http.ResponseWriter, r *http.Request) {
	generos, err := genero.GetGeneros(app.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(generos)
}

func (app *App) GetLibrosHandler(w http.ResponseWriter, r *http.Request) {
	libros, err := libro.GetLibros(app.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(libros)
}

func (app *App) GetEditorialesHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Función no implementada", http.StatusNotImplemented)
}

func (app *App) GetAutoresHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Función no implementada", http.StatusNotImplemented)
}
