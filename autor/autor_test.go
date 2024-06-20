package autor

import (
	"testing"
)

func TestNewAutor(t *testing.T) {
	autor := NewAutor(1, "Gabriel García Márquez", "Colombiano")
	if autor.GetID() != 1 {
		t.Errorf("Expected ID 1, but got %d", autor.GetID())
	}
	if autor.GetNombre() != "Gabriel García Márquez" {
		t.Errorf("Expected Nombre 'Gabriel García Márquez', but got %s", autor.GetNombre())
	}
	if autor.GetNacionalidad() != "Colombiano" {
		t.Errorf("Expected Nacionalidad 'Colombiano', but got %s", autor.GetNacionalidad())
	}
}
