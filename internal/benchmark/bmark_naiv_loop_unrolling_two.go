package benchmark

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkNaivLoopUnrollingTwo(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		naivLoopUnrollingTwo(matricesA[i], matricesB[i])
	}
}

func naivLoopUnrollingTwo(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.CronometrarMultiplicacionMatrices(modelos.NAIV_LOOP_UNROLLING_TWO, len(matrizA.Datos))()

	multiplicacion_matrices.NaivLoopUnrollingTwo(matrizA.Datos, matrizB.Datos)
}
