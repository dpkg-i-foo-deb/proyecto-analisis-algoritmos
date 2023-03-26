package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/utilidades"
	"sync"
)

func BmarkWinogradOriginal(matricesA []modelos.Matriz, matricesB []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matricesA {
		wg.Add(1)
		go winogradOriginal(matricesA[i], matricesB[i], wg)
	}
}

func winogradOriginal(matrizA modelos.Matriz, matrizB modelos.Matriz, wg *sync.WaitGroup) {
	defer utilidades.MedirTiempo(modelos.WINOGRAD_ORIGINAL, len(matrizA.Datos))()
	defer wg.Done()

	algoritmos.WinogradOriginal(matrizA.Datos, matrizB.Datos)
}
