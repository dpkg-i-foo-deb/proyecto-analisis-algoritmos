package test

import (
	"generador/pkg/algoritmos"
	"testing"
)

func TestObtenerMinimo(t *testing.T) {

	x := 1
	y := 2

	esperado := 1

	resultado := algoritmos.ObtenerMinimo(x, y)

	if resultado != esperado {
		t.Error("Obtener mínimo ha fallado")
	}
}
