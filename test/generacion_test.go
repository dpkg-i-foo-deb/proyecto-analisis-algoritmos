package test

import (
	"generador/utilidades"
	"math"
	"os"
	"strconv"
	"testing"
)

const CASOS_PRUEBA = 12

func TestGeneracionLecturaMatrices(t *testing.T) {

	defer eliminarMatrices()

	utilidades.GenerarMatrices()

	matricesA, matricesB := utilidades.LeerMatrices()

	if len(matricesA) != 12 {
		t.Error("Las matrices a no son suficientes, fueron generadas ", len(matricesA))
	}

	if len(matricesB) != 12 {
		t.Error("Las matrices b no son suficientes, fueron generadas ", len(matricesB))
	}
}

func eliminarMatrices() {
	for i := 1; i <= CASOS_PRUEBA; i++ {
		cantidad := int64(math.Pow(2.0, float64(i)))

		err := os.Remove("matriz_" + strconv.FormatInt(cantidad, 10) + "x" + strconv.FormatInt(cantidad, 10) + "_a" + ".json")

		utilidades.VerificarError(err)

	}

	for i := 1; i <= CASOS_PRUEBA; i++ {
		cantidad := int64(math.Pow(2.0, float64(i)))

		err := os.Remove("matriz_" + strconv.FormatInt(cantidad, 10) + "x" + strconv.FormatInt(cantidad, 10) + "_b" + ".json")

		utilidades.VerificarError(err)
	}
}
