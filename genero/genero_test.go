package genero

import (
	"testing"
)

func TestNewGenero(t *testing.T) {
	genero := NewGenero(1, "Fantasía")
	if genero.GetID() != 1 {
		t.Errorf("Expected ID 1, but got %d", genero.GetID())
	}
	if genero.GetNombre() != "Fantasía" {
		t.Errorf("Expected Nombre 'Fantasía', but got %s", genero.GetNombre())
	}
}
