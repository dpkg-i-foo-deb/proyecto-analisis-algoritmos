package test

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"reflect"
	"testing"
)

func TestRestarMatrices(t *testing.T) {
	a := [][]int{
		{1, 1},
		{1, 1},
	}

	b := [][]int{
		{1, 1},
		{1, 1},
	}

	esperado := [][]int{
		{0, 0},
		{0, 0},
	}

	resultado := multiplicacion_matrices.RestarMatrices(a, b)

	if !reflect.DeepEqual(resultado, esperado) {
		t.Error("Sumar matrices ha fallado")
	}
}
