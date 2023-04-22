package test

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"reflect"
	"testing"
)

func TestNaivKahan(t *testing.T) {
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

	resultado := multiplicacion_matrices.NaivKahan(a, b)

	if !reflect.DeepEqual(resultado, esperado) {
		t.Error("NaivKahan ha fallado")
	}

}
