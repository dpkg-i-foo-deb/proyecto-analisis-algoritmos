package test

import (
	"fmt"
	"generador/algoritmos"
	"reflect"
	"testing"
)

// Implementado por ChatGPT, un modelo de lenguaje entrenado por OpenAI
// Fuente: https://openaidialog.ai
func TestStrassenWinograd(t *testing.T) {
	a := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	b := [][]int{
		{17, 18, 19, 20},
		{21, 22, 23, 24},
		{25, 26, 27, 28},
		{29, 30, 31, 32},
	}

	esperado := [][]int{
		{250, 260, 270, 280},
		{618, 644, 670, 696},
		{986, 1028, 1070, 1112},
		{1354, 1412, 1470, 1528},
	}

	resultado := algoritmos.StrassenWinograd(a, b)

	fmt.Println(resultado)

	if !reflect.DeepEqual(resultado, esperado) {
		t.Error("StrassenWinograd ha fallado")
	}
}
