package test

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"reflect"
	"testing"
)

func TestWinogradOriginal(t *testing.T) {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{3, 2, 4},
	}

	b := [][]int{
		{7, 8, 9},
		{9, 6, 4},
		{11, 12, 2},
	}

	esperado := [][]int{
		{58, 56, 23},
		{139, 134, 68},
		{83, 84, 43},
	}

	resultado := multiplicacion_matrices.WinogradOriginal(a, b)

	if !reflect.DeepEqual(resultado, esperado) {
		t.Error("WinogradOriginal ha fallado")
	}
}
