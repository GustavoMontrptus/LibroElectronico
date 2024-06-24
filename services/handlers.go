package services

import (
	"net/http"
)

// Ejemplo de función manejadora adicional
func MiHandlerAdicional(w http.ResponseWriter, r *http.Request) {
	// Lógica de la función
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hola desde el manejador adicional"))
}
