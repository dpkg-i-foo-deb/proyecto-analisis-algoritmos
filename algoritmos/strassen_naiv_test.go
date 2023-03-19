package algoritmos_test

import (
	"generador/algoritmos"
	"reflect"
	"testing"
)

func TestStrassenNaiv(t *testing.T) {
	a := [][]int{
		{1, 2},
		{3, 4},
	}

	b := [][]int{
		{5, 6},
		{7, 8},
	}

	esperado := [][]int{
		{19, 22},
		{43, 50},
	}

	resultado := algoritmos.StrassenNaiv(a, b)

	if !reflect.DeepEqual(resultado, esperado) {
		t.Error("StrassenNaiv ha fallado")
	}

}
