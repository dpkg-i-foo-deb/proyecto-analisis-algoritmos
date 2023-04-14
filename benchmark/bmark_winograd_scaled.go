package benchmark

import (
	"generador/pkg/algoritmos"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkWinogradScaled(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		winogradScaled(matricesA[i], matricesB[i])
	}
}

func winogradScaled(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.WINOGRAD_SCALED, len(matrizA.Datos))()

	algoritmos.WinogradScaled(matrizA.Datos, matrizB.Datos)
}
