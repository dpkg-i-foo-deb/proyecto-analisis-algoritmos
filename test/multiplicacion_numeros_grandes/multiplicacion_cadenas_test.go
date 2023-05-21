package test

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"testing"
)

func TestMultiplicacionCadenas(t *testing.T) {
	n1 := "419236001010095628945783"
	n2 := "545656784340959960400090"

	esperado := "228758968191132222725707559027530174033598320470"

	resultado := multiplicacion_numeros_grandes.MultiplicarUsandoCadenas(n1, n2)

	if resultado != esperado {
		t.Error("La multiplicación usando cadenas ha fallado")
	}
}
