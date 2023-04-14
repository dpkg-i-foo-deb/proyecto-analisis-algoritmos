package benchmark

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func Bmark_iii_sequential_block(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		iii_sequential_block(matricesA[i], matricesB[i])
	}
}

func iii_sequential_block(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.CronometrarMultiplicacionMatrices(modelos.III_SEQUENTIAL_BLOCK, len(matrizA.Datos))()

	multiplicacion_matrices.III_SequentialBlock(matrizA.Datos, matrizB.Datos)
}
