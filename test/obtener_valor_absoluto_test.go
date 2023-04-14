package test

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"testing"
)

func TestObtenerValorAbsoluto(t *testing.T) {
	a := -5
	esperado := 5
	resultado := multiplicacion_matrices.ObtenerValorAbsoluto(a)

	if resultado != esperado {
		t.Error("ObtenerValorAbsoluto ha fallado")
	}
}
