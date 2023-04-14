package test

import (
	"generador/pkg/algoritmos"
	"testing"
)

func TestNormInf(t *testing.T) {

	A := [][]int{
		{1, -2, 3},
		{-4, 5, -6},
		{7, -8, 9},
	}

	esperado := 24

	result := algoritmos.NormInf(A)

	if result != esperado {
		t.Error("NormInf ha fallado")
	}
}
