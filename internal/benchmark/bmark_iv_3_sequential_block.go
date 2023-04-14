package benchmark

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func Bmark_iv_3_sequential_block(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		iv_3_sequential_block(matricesA[i], matricesB[i])
	}
}

func iv_3_sequential_block(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.CronometrarMultiplicacionMatrices(modelos.IV_3_SEQUENTIAL_BLOCK, len(matrizA.Datos))()

	multiplicacion_matrices.IV_3_SequentialBlock(matrizA.Datos, matrizB.Datos)
}
