package test

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"reflect"
	"testing"
)

func TestSumarMatrices(t *testing.T) {
	a := [][]int{
		{1, 1},
		{1, 1},
	}

	b := [][]int{
		{1, 1},
		{1, 1},
	}

	esperado := [][]int{
		{2, 2},
		{2, 2},
	}

	resultado := multiplicacion_matrices.SumarMatrices(a, b)

	if !reflect.DeepEqual(resultado, esperado) {
		t.Error("Sumar matrices ha fallado")
	}
}
