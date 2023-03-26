package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/tiempo"
	"sync"
)

func Bmark_V_4_parallel_block(matricesA []modelos.Matriz, matricesB []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matricesA {
		wg.Add(1)
		go v_4_parallel_block(matricesA[i], matricesB[i], wg)
	}
}

func v_4_parallel_block(matrizA modelos.Matriz, matrizB modelos.Matriz, wg *sync.WaitGroup) {
	defer tiempo.MedirTiempo(modelos.V_4_PARALLEL_BLOCK, len(matrizA.Datos))()
	defer wg.Done()

	algoritmos.V_4_ParallelBlock(matrizA.Datos, matrizB.Datos)
}
