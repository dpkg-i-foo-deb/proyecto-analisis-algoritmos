package test

import (
	"fmt"
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/utilidades"
	"testing"
)

func TestMultiplicacionInglesaIterativa(t *testing.T) {

	n1 := utilidades.FormatearCadenaASlice("12")

	n2 := utilidades.FormatearCadenaASlice("34")

	resultado := make([]int, len(n1)+len(n2))

	esperado := "408"

	resultadoCadena := utilidades.FormatearSliceACadena(multiplicacion_numeros_grandes.
		MultiplicacionInglesaIterativa(n1, n2, resultado))

	if resultadoCadena != esperado {
		t.Error("La multiplicación inglesa iterativa ha fallado")
	}

}

func TestMultiplicacionInglesaRecursiva(t *testing.T) {

	n1 := utilidades.FormatearCadenaASlice("419236001010095628945783")

	n2 := utilidades.FormatearCadenaASlice("545656784340959960400090")

	resultado := make([]int, len(n1)+len(n2))

	esperado := "228758968191132222725707559027530174033598320470"

	resultado = multiplicacion_numeros_grandes.MultiplicacionInglesaRecursiva(n1, n2, resultado, 0, 0)

	fmt.Printf("Resultado: %v", resultado)

	resultadoCadena := utilidades.FormatearSliceACadena(resultado)

	if resultadoCadena != esperado {
		t.Error("La multiplicación inglesa recursiva ha fallado")
	}

}