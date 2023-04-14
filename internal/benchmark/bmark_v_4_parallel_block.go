package benchmark

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func Bmark_V_4_parallel_block(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		v_4_parallel_block(matricesA[i], matricesB[i])
	}
}

func v_4_parallel_block(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.V_4_PARALLEL_BLOCK, len(matrizA.Datos))()

	multiplicacion_matrices.V_4_ParallelBlock(matrizA.Datos, matrizB.Datos)
}
