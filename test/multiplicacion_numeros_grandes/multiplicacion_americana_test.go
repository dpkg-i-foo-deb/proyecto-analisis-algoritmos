package test

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/utilidades"
	"testing"
)

func TestMultiplicacionAmericanaIterativa(t *testing.T) {

	n1 := utilidades.FormatearEnteroASlice("419236001010095628945783")

	n2 := utilidades.FormatearEnteroASlice("545656784340959960400090")

	esperado := "228758968191132222725707559027530174033598320470"

	resultado := utilidades.FormatearSliceAEntero(multiplicacion_numeros_grandes.
		MultiplicacionAmericanaIterativa(n1, n2))

	if resultado != esperado {
		t.Error("La multiplicación americana iterativa ha fallado")
	}

}
