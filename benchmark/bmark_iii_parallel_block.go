package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/utilidades"
	"sync"
)

func Bmark_iii_parallel_block(matricesA []modelos.Matriz, matricesB []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matricesA {
		wg.Add(1)
		go iii_parallel_block(matricesA[i], matricesB[i], wg)
	}
}

func iii_parallel_block(matrizA modelos.Matriz, matrizB modelos.Matriz, wg *sync.WaitGroup) {
	defer utilidades.MedirTiempo(modelos.III_PARALLEL_BLOCK, len(matrizA.Datos))()
	defer wg.Done()

	algoritmos.III_ParallelBlock(matrizA.Datos, matrizB.Datos)
}
