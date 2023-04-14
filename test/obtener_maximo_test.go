package test

import (
	"generador/pkg/algoritmos"
	"testing"
)

func TestObtenerMaximo(t *testing.T) {

	a := 1
	b := 2

	esperado := 2

	resultado := algoritmos.ObtenerMaximo(a, b)

	if resultado != esperado {
		t.Error("Obtener máximo ha fallado")
	}
}
