package test

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"reflect"
	"testing"
)

func TestWinogradScaled(t *testing.T) {

	A := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	B := [][]int{
		{1, 2, 4},
		{3, 4, 5},
		{5, 6, 9},
	}

	esperado := [][]int{
		{22, 28, 41},
		{49, 64, 95},
		{76, 100, 149},
	}

	resultado := multiplicacion_matrices.WinogradScaled(A, B)
	if !reflect.DeepEqual(resultado, esperado) {
		t.Error("WinogradScaled ha fallado")
	}
}
