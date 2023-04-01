package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/tiempo"
)

func BmarkNaivLoopUnrollingFour(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		naivLoopUnrollingFour(matricesA[i], matricesB[i])
	}
}

func naivLoopUnrollingFour(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.NAIV_LOOP_UNROLLING_FOUR, len(matrizA.Datos))()

	algoritmos.NaivLoopUnrollingFour(matrizA.Datos, matrizB.Datos)
}
