package bmark_multiplicacion_matrices

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkNaivLoopUnrollingFour(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		naivLoopUnrollingFour(matricesA[i], matricesB[i])
	}
}

func naivLoopUnrollingFour(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.CronometrarMultiplicacionMatrices(modelos.NAIV_LOOP_UNROLLING_FOUR, len(matrizA.Datos))()

	multiplicacion_matrices.NaivLoopUnrollingFour(matrizA.Datos, matrizB.Datos)
}
