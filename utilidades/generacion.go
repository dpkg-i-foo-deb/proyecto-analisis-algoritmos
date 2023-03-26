package utilidades

import (
	"math"
	"math/rand"
	"os"
	"strconv"

	"time"
)

const CASOS_PRUEBA = 12

func GenerarMatrices() {

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= CASOS_PRUEBA; i++ {
		cantidad := int64(math.Pow(2.0, float64(i)))

		matriz := generateMatrix(int(cantidad), int(cantidad))

		archivo, err := os.Create("matriz_" + strconv.FormatInt(cantidad, 10) + "x" + strconv.FormatInt(cantidad, 10) + ".json")

		VerificarError(err)

		writeMatrix(archivo, matriz)
	}

}
