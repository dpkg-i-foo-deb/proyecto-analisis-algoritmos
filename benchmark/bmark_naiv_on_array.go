package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/tiempo"
	"sync"
)

func BmarkNaivOnArray(matricesA []modelos.Matriz, matricesB []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matricesA {
		wg.Add(1)
		go naivOnArray(matricesA[i], matricesB[i], wg)
	}
}

func naivOnArray(matrizA modelos.Matriz, matrizB modelos.Matriz, wg *sync.WaitGroup) {
	defer tiempo.MedirTiempo(modelos.NAIV_ON_ARRAY, len(matrizA.Datos))()
	defer wg.Done()

	algoritmos.NaivOnArray(matrizA.Datos, matrizB.Datos)
}
