package utilidades

import (
	"generador/modelos"
	"math"
	"os"
	"strconv"
)

func LeerMatrices() []modelos.Matriz {

	var archivos []os.File
	var matrices []modelos.Matriz

	for i := 1; i < CASOS_PRUEBA; i++ {
		cantidad := int64(math.Pow(2.0, float64(i)))

		archivo, err := os.Open(("matriz_" + strconv.FormatInt(cantidad, 10) + "x" + strconv.FormatInt(cantidad, 10) + ".json"))

		VerificarError(err)

		archivos = append(archivos, *archivo)
	}

	for i := range archivos {
		matrices = append(matrices, readMatrix(&archivos[i]))
	}

	return matrices
}
