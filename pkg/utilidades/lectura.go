package utilidades

import (
	"generador/pkg/modelos"
	"math"
	"os"
	"strconv"
)

func LeerMatrices() ([]modelos.Matriz, []modelos.Matriz) {

	var archivosA []os.File
	var archivosB []os.File
	var matricesA []modelos.Matriz
	var matricesB []modelos.Matriz

	for i := 1; i <= CASOS_PRUEBA_MATRICES; i++ {
		cantidad := int64(math.Pow(2.0, float64(i)))

		archivo, err := os.Open(("matriz_" + strconv.FormatInt(cantidad, 10) + "x" + strconv.FormatInt(cantidad, 10) + "_a" + ".json"))

		VerificarError(err)

		archivosA = append(archivosA, *archivo)
	}

	for i := 1; i <= CASOS_PRUEBA_MATRICES; i++ {
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

func LeerNumeros() ([]modelos.NumeroGrande, []modelos.NumeroGrande) {
	var archivosA []os.File
	var archivosB []os.File
	var numerosA []modelos.NumeroGrande
	var numerosB []modelos.NumeroGrande

	n := 3

	for i := 0; i < CASOS_PRUEBA_NUMEROS; i++ {
		cantidad := int((math.Pow(2.0, float64(n))))

		archivo, err := os.Open("numero_" + strconv.FormatInt(int64(cantidad), 10) + "_a" + ".json")

		VerificarError(err)

		archivosA = append(archivosA, *archivo)

		archivo2, err := os.Open("numero_" + strconv.FormatInt(int64(cantidad), 10) + "_b" + ".json")

		VerificarError(err)

		archivosB = append(archivosB, *archivo2)

		n++
	}

	for i := range archivosA {
		numerosA = append(numerosA, leerNumero(&archivosA[i]))
	}

	for i := range archivosB {
		numerosB = append(numerosB, leerNumero(&archivosB[i]))
	}

	return numerosA, numerosB
}

func LeerResultados() []modelos.ResultadoMultiplicacionMatrices {
	archivo, err := os.Open("resultados.json")

	VerificarError(err)

	resultados := readResult(archivo)

	return resultados
}
