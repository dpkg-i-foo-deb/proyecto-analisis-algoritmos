package test

import (
	"generador/pkg/algoritmos"
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

	resultado := algoritmos.SumarMatrices(a, b)

	if !reflect.DeepEqual(resultado, esperado) {
		t.Error("Sumar matrices ha fallado")
	}
}
