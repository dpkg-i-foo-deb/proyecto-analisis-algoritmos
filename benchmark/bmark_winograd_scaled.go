package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/tiempo"
	"sync"
)

func BmarkWinogradScaled(matricesA []modelos.Matriz, matricesB []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matricesA {
		wg.Add(1)
		go winogradScaled(matricesA[i], matricesB[i], wg)
	}
}

func winogradScaled(matrizA modelos.Matriz, matrizB modelos.Matriz, wg *sync.WaitGroup) {
	defer tiempo.MedirTiempo(modelos.WINOGRAD_SCALED, len(matrizA.Datos))()
	defer wg.Done()

	algoritmos.WinogradScaled(matrizA.Datos, matrizB.Datos)
}
