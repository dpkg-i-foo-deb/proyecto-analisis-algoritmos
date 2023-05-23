package test

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/utilidades"
	"testing"
)

func TestMultiplicacionDivideVenceras1(t *testing.T) {

	n1 := utilidades.FormatearCadenaASlice("847")

	n2 := utilidades.FormatearCadenaASlice("69")

	resultado := make([]int, len(n1)+len(n2)*2)

	esperado := "58443"

	resultadoCadena := utilidades.FormatearSliceACadena(multiplicacion_numeros_grandes.
		MultiplicacionDivideVenceras1(n1, n2, resultado))

	if resultadoCadena != esperado {
		t.Error("La multiplicación divide y vencerás 1 ha fallado")
	}

}
