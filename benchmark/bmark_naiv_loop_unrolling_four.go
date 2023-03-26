package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/utilidades"
	"sync"
)

func BmarkNaivLoopUnrollingFour(matricesA []modelos.Matriz, matricesB []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matricesA {
		wg.Add(1)
		go naivLoopUnrollingFour(matricesA[i], matricesB[i], wg)
	}
}

func naivLoopUnrollingFour(matrizA modelos.Matriz, matrizB modelos.Matriz, wg *sync.WaitGroup) {
	defer utilidades.MedirTiempo(modelos.NAIV_LOOP_UNROLLING_FOUR, len(matrizA.Datos))()
	defer wg.Done()

	algoritmos.NaivLoopUnrollingFour(matrizA.Datos, matrizB.Datos)
}
