package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/tiempo"
	"sync"
)

func Bmark_iv_3_sequential_block(matricesA []modelos.Matriz, matricesB []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matricesA {
		wg.Add(1)
		go iv_3_sequential_block(matricesA[i], matricesB[i], wg)
	}
}

func iv_3_sequential_block(matrizA modelos.Matriz, matrizB modelos.Matriz, wg *sync.WaitGroup) {
	defer tiempo.MedirTiempo(modelos.IV_3_SEQUENTIAL_BLOCK, len(matrizA.Datos))()
	defer wg.Done()

	algoritmos.IV_3_SequentialBlock(matrizA.Datos, matrizB.Datos)
}
