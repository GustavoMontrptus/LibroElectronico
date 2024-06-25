package services

import (
	"Proyecto/autor"
	"Proyecto/editorial"
	"Proyecto/genero"
	"Proyecto/libro"
	"database/sql"
	"encoding/json"
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
