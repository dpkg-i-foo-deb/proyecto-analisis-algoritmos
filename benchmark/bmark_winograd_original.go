package benchmark

import (
	"generador/algoritmos"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkWinogradOriginal(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		winogradOriginal(matricesA[i], matricesB[i])
	}
}

func winogradOriginal(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.WINOGRAD_ORIGINAL, len(matrizA.Datos))()

	algoritmos.WinogradOriginal(matrizA.Datos, matrizB.Datos)
}
