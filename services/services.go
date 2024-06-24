package services

import (
	"Proyecto/autor"
	"Proyecto/editorial"
	"Proyecto/genero"
	"Proyecto/libro"
	"Proyecto/usuario"
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
)

type App struct {
	DB *sql.DB
}

func (app *App) GetAutoresHandler(w http.ResponseWriter, r *http.Request) {
	autores, err := autor.GetAutores(app.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(autores)
}

func (app *App) GetEditorialesHandler(w http.ResponseWriter, r *http.Request) {
	editoriales, err := editorial.GetEditoriales(app.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(editoriales)
}

func (app *App) GetGenerosHandler(w http.ResponseWriter, r *http.Request) {
	generos, err := genero.GetGeneros(app.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(generos)
}

func (app *App) GetLibrosHandler(w http.ResponseWriter, r *http.Request) {
	libros, err := libro.GetLibros(app.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(libros)
}

func (app *App) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "templates/login.html")
		return
	}

	correo := r.FormValue("correo")
	contrasena := r.FormValue("contrasena")

	if !usuario.ValidarCorreo(correo) || !usuario.ValidarContrasena(contrasena) {
		http.Error(w, "Correo o contraseña inválidos", http.StatusBadRequest)
		return
	}

	_, autenticado := usuario.AutenticarUsuario(app.DB, correo, contrasena)
	if autenticado {
		http.Redirect(w, r, "/seleccion-genero", http.StatusSeeOther)
	} else {
		err := usuario.CrearCliente(app.DB, correo, contrasena)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/seleccion-genero", http.StatusSeeOther)
	}
}

func (app *App) SeleccionGeneroHandler(w http.ResponseWriter, r *http.Request) {
	generos, err := genero.GetGeneros(app.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Generos []genero.Genero
	}{
		Generos: generos,
	}

	tmpl := template.Must(template.ParseFiles("templates/seleccion_genero.html"))
	tmpl.Execute(w, data)
}

func (app *App) LibrosHandler(w http.ResponseWriter, r *http.Request) {
	generos, err := genero.GetGeneros(app.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	libros, err := libro.GetLibros(app.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Generos []genero.Genero
		Libros  []libro.Libro
	}{
		Generos: generos,
		Libros:  libros,
	}

	tmpl := template.Must(template.ParseFiles("templates/libros.html"))
	tmpl.Execute(w, data)
}
