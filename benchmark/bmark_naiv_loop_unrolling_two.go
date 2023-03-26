package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/tiempo"
	"sync"
)

func BmarkNaivLoopUnrollingTwo(matricesA []modelos.Matriz, matricesB []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matricesA {
		wg.Add(1)
		go naivLoopUnrollingTwo(matricesA[i], matricesB[i], wg)
	}
}

func naivLoopUnrollingTwo(matrizA modelos.Matriz, matrizB modelos.Matriz, wg *sync.WaitGroup) {
	defer tiempo.MedirTiempo(modelos.NAIV_LOOP_UNROLLING_TWO, len(matrizA.Datos))()
	defer wg.Done()

	algoritmos.NaivLoopUnrollingTwo(matrizA.Datos, matrizB.Datos)
}
