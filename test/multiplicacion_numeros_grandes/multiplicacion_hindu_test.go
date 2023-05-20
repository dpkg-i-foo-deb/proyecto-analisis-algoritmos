package test

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/utilidades"
	"testing"
)

func TestMultiplicacionHinduIterativa(t *testing.T) {

	n1 := utilidades.FormatearCadenaASlice("419236001010095628945783")

	n2 := utilidades.FormatearCadenaASlice("545656784340959960400090")

	resultado := make([]int, len(n1)+len(n2))

	esperado := "228758968191132222725707559027530174033598320470"

	resultadoCadena := utilidades.FormatearSliceACadena(multiplicacion_numeros_grandes.
		MultiplicacionHinduIterativa(n1, n2, resultado))

	if resultadoCadena != esperado {
		t.Error("La multiplicación Hindu iterativa ha fallado")
	}

}

func TestMultiplicacionHinduRecursiva(t *testing.T) {

	n1 := utilidades.FormatearCadenaASlice("419236001010095628945783")

	n2 := utilidades.FormatearCadenaASlice("545656784340959960400090")

	resultado := make([]int, len(n1)+len(n2))

	esperado := "228758968191132222725707559027530174033598320470"

	resultado = multiplicacion_numeros_grandes.MultiplicacionHinduRecursiva(n1, n2, resultado, 0, 0)

	resultadoCadena := utilidades.FormatearSliceACadena(resultado)

	if resultadoCadena != esperado {
		t.Error("La multiplicación hindu recursiva ha fallado")
	}

}
