package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/utilidades"
	"sync"
)

func BmarkStrassenWinograd(matricesA []modelos.Matriz, matricesB []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matricesA {
		wg.Add(1)
		go strassenWinograd(matricesA[i], matricesB[i], wg)
	}
}

func strassenWinograd(matrizA modelos.Matriz, matrizB modelos.Matriz, wg *sync.WaitGroup) {
	defer utilidades.MedirTiempo(modelos.STRASSEN_WINOGRAD, len(matrizA.Datos))()
	defer wg.Done()

	algoritmos.StrassenWinograd(matrizA.Datos, matrizB.Datos)
}
