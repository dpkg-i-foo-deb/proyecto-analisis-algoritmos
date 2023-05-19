package test

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/utilidades"
	"testing"
)

func TestMultiplicacionRusaIterativa(t *testing.T) {

	n1 := utilidades.FormatearCadenaASlice("12908")

	n2 := utilidades.FormatearCadenaASlice("34798")

	resultado := make([]int, len(n1)+len(n2))

	esperado := "449172584"

	resultadoCadena := utilidades.FormatearSliceACadena(multiplicacion_numeros_grandes.
		MultiplicacionRusaIterativa(n1, n2, resultado))

	if resultadoCadena != esperado {
		t.Error("La multiplicación Rusa iterativa ha fallado")
	}

}

func TestMultiplicacionRusaRecursiva(t *testing.T) {

	n1 := utilidades.FormatearCadenaASlice("12908")

	n2 := utilidades.FormatearCadenaASlice("34798")

	resultado := make([]int, len(n1)+len(n2))

	esperado := "449172584"

	resultado = multiplicacion_numeros_grandes.MultiplicacionRusaRecursiva(n1, n2, resultado)

	resultadoCadena := utilidades.FormatearSliceACadena(resultado)

	if resultadoCadena != esperado {
		t.Error("La multiplicación inglesa recursiva ha fallado")
	}

}
