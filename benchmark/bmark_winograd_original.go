package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/tiempo"
	"sync"
)

func BmarkWinogradOriginal(matricesA []modelos.Matriz, matricesB []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matricesA {
		//wg.Add(1)
		winogradOriginal(matricesA[i], matricesB[i], wg)
	}
}

func winogradOriginal(matrizA modelos.Matriz, matrizB modelos.Matriz, wg *sync.WaitGroup) {
	defer tiempo.MedirTiempo(modelos.WINOGRAD_ORIGINAL, len(matrizA.Datos))()
	//defer wg.Done()

	algoritmos.WinogradOriginal(matrizA.Datos, matrizB.Datos)
}
