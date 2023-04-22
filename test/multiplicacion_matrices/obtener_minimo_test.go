package test

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"testing"
)

func TestObtenerMinimo(t *testing.T) {

	x := 1
	y := 2

	esperado := 1

	resultado := multiplicacion_matrices.ObtenerMinimo(x, y)

	if resultado != esperado {
		t.Error("Obtener mínimo ha fallado")
	}
}
