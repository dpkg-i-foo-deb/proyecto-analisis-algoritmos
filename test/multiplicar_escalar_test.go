package test

import (
	"generador/algoritmos"
	"reflect"
	"testing"
)

func TestMultiplicarEscalar(t *testing.T) {

	A := [][]int{
		{1, 2},
		{3, 4},
	}

	Resultado := make([][]int, len(A))

	for i := range Resultado {
		Resultado[i] = make([]int, len(A[0]))
	}
	alpha := 2

	algoritmos.MultiplicarEscalar(A, Resultado, alpha)

	expected := [][]int{{2, 4}, {6, 8}}
	if !reflect.DeepEqual(Resultado, expected) {
		t.Errorf("Expected %v but got %v", expected, Resultado)
	}
}
