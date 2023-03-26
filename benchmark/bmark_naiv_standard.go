package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/tiempo"
	"sync"
)

func BmarkNaivStandard(matricesA []modelos.Matriz, matricesB []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matricesA {
		wg.Add(1)
		go naivStandard(matricesA[i], matricesB[i], wg)
	}
}

func naivStandard(matrizA modelos.Matriz, matrizB modelos.Matriz, wg *sync.WaitGroup) {
	defer tiempo.MedirTiempo(modelos.NAIV_STANDARD, len(matrizA.Datos))()
	defer wg.Done()

	algoritmos.NaivStandard(matrizA.Datos, matrizB.Datos)
}
