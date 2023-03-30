package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/tiempo"
	"sync"
)

func Bmark_iv_4_parallel_block(matricesA []modelos.Matriz, matricesB []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matricesA {
		//wg.Add(1)
		iv_4_parallel_block(matricesA[i], matricesB[i], wg)
	}
}

func iv_4_parallel_block(matrizA modelos.Matriz, matrizB modelos.Matriz, wg *sync.WaitGroup) {
	defer tiempo.MedirTiempo(modelos.IV_4_PARALLEL_BLOCK, len(matrizA.Datos))()
	//defer wg.Done()

	algoritmos.IV_4_Pararell_Block(matrizA.Datos, matrizB.Datos)
}
