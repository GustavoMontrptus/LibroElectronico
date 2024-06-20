package libro

import (
	"Proyecto/autor"
	"Proyecto/editorial"
	"Proyecto/genero"
	"testing"
)

func TestMostrarLibrosPorGenero(t *testing.T) {
	autores := []autor.Autor{
		{ID: 1, Nombre: "J.R.R. Tolkien", Nacionalidad: "Inglés"},
		{ID: 2, Nombre: "Gabriel García Márquez", Nacionalidad: "Colombiano"},
	}

	editoriales := []editorial.Editorial{
		{ID: 1, Nombre: "Allen & Unwin"},
		{ID: 2, Nombre: "Editorial Sudamericana"},
	}

	generos := []genero.Genero{
		{ID: 1, Nombre: "Fantasía"},
		{ID: 2, Nombre: "Realismo Mágico"},
	}

	tests := []struct {
		generoID int
		expected bool
	}{
		{1, true},
		{2, true},
		{99, false}, // ID de género inexistente
	}

	for _, test := range tests {
		result := MostrarLibrosPorGenero(test.generoID, autores, editoriales, generos)
		if result != test.expected {
			t.Errorf("MostrarLibrosPorGenero(%d) = %v; want %v", test.generoID, result, test.expected)
		}
	}
}
