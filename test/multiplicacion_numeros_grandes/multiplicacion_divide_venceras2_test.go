package test

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/utilidades"
	"testing"
)

func TestMultiplicacionDivideVenceras2(t *testing.T) {
	n1 := utilidades.FormatearCadenaASlice("10")

	n2 := utilidades.FormatearCadenaASlice("8")

	resultado := make([]int, len(n1)+len(n2)*2)

	esperado := "80"

	resultadoCadena := utilidades.FormatearSliceACadena(multiplicacion_numeros_grandes.
		MultiplicacionDivideVenceras2(n1, n2, resultado))

	if resultadoCadena != esperado {
		t.Error("La multiplicación americana iterativa ha fallado")
	}
}
