package test

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"testing"
)

func TestNormInf(t *testing.T) {

	A := [][]int{
		{1, -2, 3},
		{-4, 5, -6},
		{7, -8, 9},
	}

	esperado := 24

	result := multiplicacion_matrices.NormInf(A)

	if result != esperado {
		t.Error("NormInf ha fallado")
	}
}
