package benchmark

import (
	"generador/algoritmos"
	"generador/pkg/modelos"
	"generador/tiempo"
)

func Bmark_V_4_parallel_block(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		v_4_parallel_block(matricesA[i], matricesB[i])
	}
}

func v_4_parallel_block(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.V_4_PARALLEL_BLOCK, len(matrizA.Datos))()

	algoritmos.V_4_ParallelBlock(matrizA.Datos, matrizB.Datos)
}
