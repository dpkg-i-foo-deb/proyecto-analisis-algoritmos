package algoritmos_test

import (
	"generador/algoritmos"
	"generador/utilidades"
	"reflect"
	"testing"
)

func TestNaivStandard(t *testing.T) {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	b := [][]int{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	esperado := [][]int{
		{58, 64},
		{139, 154},
	}

	resultado := algoritmos.NaivStandard(a, b)

	if reflect.DeepEqual(resultado, esperado) {
		utilidades.ImprimirMatriz(resultado)
		utilidades.ImprimirMatriz(esperado)
		t.Error("NaivStandard ha fallado")
	}
}
