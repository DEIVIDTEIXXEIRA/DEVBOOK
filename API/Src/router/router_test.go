package router

import "testing"

func TestConectar(t *testing.T) {
	route := Gerar()

	if route == nil {
		t.Fatal("O router não deve ser nulo")
	}
}