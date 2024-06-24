package main

import (
	"Proyecto/services"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func initDB() *sql.DB {
	db, err := sql.Open("mysql", "root:camila262004@tcp(127.0.0.1:3306)/biblioteca_go")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database successfully!")
	return db
}

func main() {
	db := initDB()
	defer db.Close()

	app := &services.App{DB: db}

	http.HandleFunc("/login", app.LoginHandler)
	http.HandleFunc("/libros", app.LibrosHandler)
	http.HandleFunc("/api/autores", app.GetAutoresHandler)
	http.HandleFunc("/api/editoriales", app.GetEditorialesHandler)
	http.HandleFunc("/api/generos", app.GetGenerosHandler)
	http.HandleFunc("/api/libros", app.GetLibrosHandler)

	// Servir index.html como página principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "templates/index.html")
	})

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
