package test

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/modelos"
	"generador/pkg/utilidades"
	"testing"
)

func TestMultiplicacionAmericanaIterativa(t *testing.T) {

	n1 := utilidades.FormatearCadenaASlice("419236001010095628945783")

	n2 := utilidades.FormatearCadenaASlice("545656784340959960400090")

	resultado := make([]int, len(n1)+len(n2)*2)

	esperado := "228758968191132222725707559027530174033598320470"

	resultadoCadena := utilidades.FormatearSliceACadena(multiplicacion_numeros_grandes.
		MultiplicacionAmericanaIterativa(n1, n2, resultado))

	if resultadoCadena != esperado {
		t.Error("La multiplicación americana iterativa ha fallado")
	}

}

func TestMultiplicacionAmericanaRecursiva(t *testing.T) {
	n1 := utilidades.FormatearCadenaASlice("419236001010095628945783")

	n2 := utilidades.FormatearCadenaASlice("545656784340959960400090")

	resultado := make([]int, len(n1)+len(n2)*2)

	esperado := "228758968191132222725707559027530174033598320470"

	resultadoCadena := utilidades.FormatearSliceACadena(
		multiplicacion_numeros_grandes.MultiplicacionAmericanaRecursiva(n1, n2, resultado, len(n2)-1, len(n1)-1, len(resultado)-1))

	if resultadoCadena != esperado {
		t.Error("La multiplicación americana iterativa ha fallado")
	}
}

func TestMultiplicacionAmericanaRecursivaEstructuras(t *testing.T) {
	n1 := utilidades.FormatearCadenaALista("419236001010095628945783")

	n2 := utilidades.FormatearCadenaALista("545656784340959960400090")

	resultado := modelos.Lista((n1.GetCantidadNodos() + n2.GetCantidadNodos()) * 2)

	esperado := "228758968191132222725707559027530174033598320470"

	resultadoCadena := utilidades.FormatearListaACadena(multiplicacion_numeros_grandes.MultiplicacionAmericanaRecursivaEstructuras(n1, n2, resultado, n2.Cantidad-1, n1.Cantidad-1, resultado.Cantidad-1))

	if resultadoCadena != esperado {
		t.Error("La multiplicación americana iterativa con estructuras ha fallado")
	}
}

func TestMultiplicacionAmericanaIterativaEstructuras(t *testing.T) {
	n1 := utilidades.FormatearCadenaALista("419236001010095628945783")

	n2 := utilidades.FormatearCadenaALista("545656784340959960400090")

	resultado := modelos.Lista((n1.GetCantidadNodos() + n2.GetCantidadNodos()) * 2)

	esperado := "228758968191132222725707559027530174033598320470"

	resultadoCadena := utilidades.FormatearListaACadena(multiplicacion_numeros_grandes.MultiplicacionAmericanaIterativaEstructuras(n1, n2, resultado))

	if resultadoCadena != esperado {
		t.Error("La multiplicación americana iterativa con estructuras ha fallado")
	}
}
