package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/tiempo"
	"sync"
)

func BmarkStrassenWinograd(matricesA []modelos.Matriz, matricesB []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matricesA {
		wg.Add(1)
		go strassenWinograd(matricesA[i], matricesB[i], wg)
	}
}

func strassenWinograd(matrizA modelos.Matriz, matrizB modelos.Matriz, wg *sync.WaitGroup) {
	defer tiempo.MedirTiempo(modelos.STRASSEN_WINOGRAD, len(matrizA.Datos))()
	defer wg.Done()

	algoritmos.StrassenWinograd(matrizA.Datos, matrizB.Datos)
}
