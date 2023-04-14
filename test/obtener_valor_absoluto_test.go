package test

import (
	"generador/pkg/algoritmos"
	"testing"
)

func TestObtenerValorAbsoluto(t *testing.T) {
	a := -5
	esperado := 5
	resultado := algoritmos.ObtenerValorAbsoluto(a)

	if resultado != esperado {
		t.Error("ObtenerValorAbsoluto ha fallado")
	}
}
