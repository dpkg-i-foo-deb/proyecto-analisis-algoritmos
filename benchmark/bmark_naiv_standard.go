package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/tiempo"
)

func BmarkNaivStandard(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		naivStandard(matricesA[i], matricesB[i])
	}
}

func naivStandard(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.NAIV_STANDARD, len(matrizA.Datos))()

	algoritmos.NaivStandard(matrizA.Datos, matrizB.Datos)
}
