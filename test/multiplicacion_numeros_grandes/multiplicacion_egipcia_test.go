package test

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/utilidades"
	"testing"
)

func TestMultiplicacionEgipciaIterativa(t *testing.T) {

	n1 := utilidades.FormatearCadenaASlice("12908")

	n2 := utilidades.FormatearCadenaASlice("34798")

	resultado := make([]int, len(n1)+len(n2))

	esperado := "449172584"

	resultadoCadena := utilidades.FormatearSliceACadena(multiplicacion_numeros_grandes.
		MultiplicacionEgipciaIterativa(n1, n2, resultado))

	if resultadoCadena != esperado {
		t.Error("La multiplicación rusa iterativa ha fallado")
	}

}
