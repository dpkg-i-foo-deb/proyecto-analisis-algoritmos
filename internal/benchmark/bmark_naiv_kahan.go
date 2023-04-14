package benchmark

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkNaivKahan(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		naivKahan(matricesA[i], matricesB[i])
	}
}

func naivKahan(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.NAIV_KAHAN, len(matrizA.Datos))()

	multiplicacion_matrices.NaivKahan(matrizA.Datos, matrizB.Datos)
}
