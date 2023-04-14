package test

import (
	"generador/pkg/multiplicacion_matrices"
	"reflect"
	"testing"
)

func TestIIISequentialBlock(t *testing.T) {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	b := [][]int{
		{10, 11, 12},
		{13, 14, 15},
		{16, 17, 18},
	}

	esperado := [][]int{
		{84, 90, 96},
		{201, 216, 231},
		{318, 342, 366},
	}

	resultado := multiplicacion_matrices.III_SequentialBlock(a, b)

	if !reflect.DeepEqual(resultado, esperado) {
		t.Error("III_Sequential Block ha fallado")
	}

}
