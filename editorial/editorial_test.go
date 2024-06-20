package editorial

import (
	"testing"
)

func TestNewEditorial(t *testing.T) {
	editorial := NewEditorial(1, "Allen & Unwin")
	if editorial.GetID() != 1 {
		t.Errorf("Expected ID 1, but got %d", editorial.GetID())
	}
	if editorial.GetNombre() != "Allen & Unwin" {
		t.Errorf("Expected Nombre 'Allen & Unwin', but got %s", editorial.GetNombre())
	}
}
