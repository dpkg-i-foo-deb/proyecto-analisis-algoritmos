package test

import (
	"generador/pkg/algoritmos"
	"reflect"
	"testing"
)

func TestV_4_ParallelBlock(t *testing.T) {
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

	resultado := algoritmos.V_4_ParallelBlock(a, b)

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("ParallelBlock ha fallado")
	}
}
