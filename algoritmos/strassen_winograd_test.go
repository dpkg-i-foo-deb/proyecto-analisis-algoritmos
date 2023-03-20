package algoritmos_test

import (
	"generador/algoritmos"
	"reflect"
	"testing"
)

func TestStrassenWinograd(t *testing.T) {
	a := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	b := [][]int{
		{9, 10},
		{11, 12},
		{13, 14},
		{15, 16},
	}

	esperado := [][]int{
		{130, 140},
		{322, 348},
	}

	resultado := algoritmos.StrassenWinograd(a, b)

	if !reflect.DeepEqual(resultado, esperado) {
		t.Error("StrassenWinograd ha fallado")
	}
}
