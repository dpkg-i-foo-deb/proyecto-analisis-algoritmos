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

	matrices := utilidades.LeerMatrices()

	if len(matrices) != 12 {
		t.Error("Las matrices no son suficientes, fueron generadas ", len(matrices))
	}
}

func eliminarMatrices() {
	for i := 1; i <= CASOS_PRUEBA; i++ {
		cantidad := int64(math.Pow(2.0, float64(i)))

		err := os.Remove("matriz_" + strconv.FormatInt(cantidad, 10) + "x" + strconv.FormatInt(cantidad, 10) + ".json")

		utilidades.VerificarError(err)

	}
}
