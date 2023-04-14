package benchmark

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func Bmark_iv_4_parallel_block(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		iv_4_parallel_block(matricesA[i], matricesB[i])
	}
}

func iv_4_parallel_block(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.CronometrarMultiplicacionMatrices(modelos.IV_4_PARALLEL_BLOCK, len(matrizA.Datos))()
	multiplicacion_matrices.IV_4_Pararell_Block(matrizA.Datos, matrizB.Datos)
}
