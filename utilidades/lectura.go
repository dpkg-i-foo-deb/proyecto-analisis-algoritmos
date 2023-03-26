package utilidades

import (
	"generador/modelos"
	"math"
	"os"
	"strconv"
)

func LeerMatrices() ([]modelos.Matriz, []modelos.Matriz) {

	var archivosA []os.File
	var archivosB []os.File
	var matricesA []modelos.Matriz
	var matricesB []modelos.Matriz

	for i := 1; i <= CASOS_PRUEBA; i++ {
		cantidad := int64(math.Pow(2.0, float64(i)))

		archivo, err := os.Open(("matriz_" + strconv.FormatInt(cantidad, 10) + "x" + strconv.FormatInt(cantidad, 10) + "_a" + ".json"))

		VerificarError(err)

		archivosA = append(archivosA, *archivo)
	}

	for i := 1; i <= CASOS_PRUEBA; i++ {
		cantidad := int64(math.Pow(2.0, float64(i)))

		archivo, err := os.Open(("matriz_" + strconv.FormatInt(cantidad, 10) + "x" + strconv.FormatInt(cantidad, 10) + "_b" + ".json"))

		VerificarError(err)

		archivosB = append(archivosB, *archivo)
	}

	for i := range archivosA {
		matricesA = append(matricesA, readMatrix(&archivosA[i]))
	}

	for i := range archivosB {
		matricesB = append(matricesB, readMatrix(&archivosB[i]))
	}

	return matricesA, matricesB
}
