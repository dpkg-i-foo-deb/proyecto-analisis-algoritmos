package test

import (
	"generador/algoritmos"
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

	resultado := algoritmos.RestarMatrices(a, b)

	if !reflect.DeepEqual(resultado, esperado) {
		t.Error("Sumar matrices ha fallado")
	}
}
