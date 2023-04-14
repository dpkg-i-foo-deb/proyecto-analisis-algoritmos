package benchmark

import (
	"generador/algoritmos"
	"generador/pkg/modelos"
	"generador/tiempo"
)

func BmarkNaivLoopUnrollingThree(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		naivLoopUnrollingThree(matricesA[i], matricesB[i])
	}
}

func naivLoopUnrollingThree(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.NAIV_LOOP_UNROLLING_THREE, len(matrizA.Datos))()

	algoritmos.NaivLoopUnrollingThree(matrizA.Datos, matrizB.Datos)
}
