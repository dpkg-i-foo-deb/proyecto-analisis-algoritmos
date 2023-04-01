package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/tiempo"
)

func BmarkNaivOnArray(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		naivOnArray(matricesA[i], matricesB[i])
	}
}

func naivOnArray(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.NAIV_ON_ARRAY, len(matrizA.Datos))()

	algoritmos.NaivOnArray(matrizA.Datos, matrizB.Datos)
}
