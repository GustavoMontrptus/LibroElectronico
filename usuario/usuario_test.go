package usuario

import (
	"testing"
)

func TestValidarCorreo(t *testing.T) {
	valid := ValidarCorreo("example@example.com")
	if !valid {
		t.Errorf("Expected true but got false")
	}

	invalid := ValidarCorreo("example.com")
	if invalid {
		t.Errorf("Expected false but got true")
	}
}

func TestValidarContrasena(t *testing.T) {
	valid := ValidarContrasena("Password123")
	if !valid {
		t.Errorf("Expected true but got false")
	}

	invalid := ValidarContrasena("password")
	if invalid {
		t.Errorf("Expected false but got true")
	}
}
