package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/utilidades"
	"sync"
)

func BmarkNaivLoopUnrollingThree(matricesA []modelos.Matriz, matricesB []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matricesA {
		wg.Add(1)
		go naivLoopUnrollingThree(matricesA[i], matricesB[i], wg)
	}
}

func naivLoopUnrollingThree(matrizA modelos.Matriz, matrizB modelos.Matriz, wg *sync.WaitGroup) {
	defer utilidades.MedirTiempo(modelos.NAIV_LOOP_UNROLLING_THREE, len(matrizA.Datos))()
	defer wg.Done()

	algoritmos.NaivLoopUnrollingThree(matrizA.Datos, matrizB.Datos)
}
