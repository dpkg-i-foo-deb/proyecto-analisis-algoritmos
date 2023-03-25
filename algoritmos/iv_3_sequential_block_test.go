package algoritmos_test

import (
	"generador/algoritmos"
	"reflect"
	"testing"
)

func Test_IV_3_Sequential_Block(t *testing.T) {

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

	resultado := algoritmos.IV_3_SequentialBlock(A, B)

	if !reflect.DeepEqual(resultado, esperado) {
		t.Error("IV.3 Sequential Block ha fallado")
	}
}
