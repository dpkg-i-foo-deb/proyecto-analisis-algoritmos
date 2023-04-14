package test

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"reflect"
	"testing"
)

func TestIV_4_ParallelBlock(t *testing.T) {
	a := [][]int{
		{1, 2},
		{4, 5},
	}

	b := [][]int{
		{10, 11},
		{13, 14},
	}

	esperado := [][]int{
		{36, 39},
		{105, 114},
	}

	resultado := multiplicacion_matrices.IV_4_Pararell_Block(a, b)

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("IV_4 ParallelBlock ha fallado")
	}
}
