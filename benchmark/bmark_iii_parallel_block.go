package benchmark

import (
	"generador/algoritmos"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func Bmark_iii_parallel_block(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		iii_parallel_block(matricesA[i], matricesB[i])
	}
}

func iii_parallel_block(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.III_PARALLEL_BLOCK, len(matrizA.Datos))()

	algoritmos.III_ParallelBlock(matrizA.Datos, matrizB.Datos)
}
