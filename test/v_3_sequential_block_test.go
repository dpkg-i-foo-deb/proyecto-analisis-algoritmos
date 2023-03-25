package test

import (
	"generador/algoritmos"
	"reflect"
	"testing"
)

func TestV_3_SequentialBlock(t *testing.T) {

	A := [][]int{
		{1, 2},
		{3, 4},
	}

	B := [][]int{
		{5, 6},
		{7, 8},
	}

	esperado := [][]int{
		{19, 22},
		{43, 50},
	}

	resultado := algoritmos.V_3_SequentialBlock(A, B)

	if !reflect.DeepEqual(resultado, esperado) {
		t.Error("V.3 Sequential Block ha fallado")
	}
}
