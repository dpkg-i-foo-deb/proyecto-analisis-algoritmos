package test

import (
	"generador/pkg/utilidades"
	"math"
	"os"
	"strconv"
	"testing"
)

const CASOS_PRUEBA_NUMEROS = 8

func TestGeneracionNumeros(t *testing.T) {
	defer eliminarNumeros()

	utilidades.GenerarNumeros()

	numerosA, numerosB := utilidades.LeerNumeros()

	if len(numerosA) != CASOS_PRUEBA_NUMEROS {
		t.Error("Los números no son suficientes")
	}

	if len(numerosB) != CASOS_PRUEBA_NUMEROS {
		t.Error("Los números no son suficientes")
	}
}

func eliminarNumeros() {
	n := 18

	for i := 0; i < CASOS_PRUEBA_NUMEROS; i++ {
		cantidad := int((math.Pow(2.0, float64(n))))

		err := os.Remove("numero_" + strconv.FormatInt(int64(cantidad), 10) + "_a" + ".json")

		utilidades.VerificarError(err)

		err = os.Remove("numero_" + strconv.FormatInt(int64(cantidad), 10) + "_b" + ".json")

		utilidades.VerificarError(err)

		n++

	}
}
