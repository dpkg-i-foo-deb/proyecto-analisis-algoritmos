package benchmark

import (
	"generador/pkg/algoritmos"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkNaivLoopUnrollingTwo(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		naivLoopUnrollingTwo(matricesA[i], matricesB[i])
	}
}

func naivLoopUnrollingTwo(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.NAIV_LOOP_UNROLLING_TWO, len(matrizA.Datos))()

	algoritmos.NaivLoopUnrollingTwo(matrizA.Datos, matrizB.Datos)
}
