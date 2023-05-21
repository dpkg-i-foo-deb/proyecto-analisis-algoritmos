package utilidades

import (
	"math"
	"math/rand"
	"os"
	"strconv"

	"time"
)

const CASOS_PRUEBA_MATRICES = 12
const CASOS_PRUEBA_NUMEROS = 8

func GenerarMatrices() {

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= CASOS_PRUEBA_MATRICES; i++ {
		cantidad := int64(math.Pow(2.0, float64(i)))

		matriz := generateMatrix(int(cantidad), int(cantidad))

		archivo, err := os.Create("matriz_" + strconv.FormatInt(cantidad, 10) + "x" + strconv.FormatInt(cantidad, 10) + "_a" + ".json")

		VerificarError(err)

		writeMatrix(archivo, matriz)
	}

	for i := 1; i <= CASOS_PRUEBA_MATRICES; i++ {
		cantidad := int64(math.Pow(2.0, float64(i)))

		matriz := generateMatrix(int(cantidad), int(cantidad))

		archivo, err := os.Create("matriz_" + strconv.FormatInt(cantidad, 10) + "x" + strconv.FormatInt(cantidad, 10) + "_b" + ".json")

		VerificarError(err)

		writeMatrix(archivo, matriz)
	}

}

func GenerarNumeros() {
	rand.Seed(time.Now().UnixNano())

	n := 20

	for i := 0; i < CASOS_PRUEBA_NUMEROS; i++ {
		cantidad := int((math.Pow(2.0, float64(n))))

		numero := generarNumeroGrande(cantidad)

		archivo, err := os.Create("numero_" + strconv.FormatInt(10, cantidad) + "_a" + ".json")

		VerificarError(err)

		escribirNumero(numero, archivo)

		n += 2
	}

	for i := 0; i < CASOS_PRUEBA_MATRICES; i++ {
		cantidad := int((math.Pow(2.0, float64(n))))

		numero := generarNumeroGrande(cantidad)

		archivo, err := os.Create("numero_" + strconv.FormatInt(10, cantidad) + "_b" + ".json")

		VerificarError(err)

		escribirNumero(numero, archivo)

		n += 2
	}
}
