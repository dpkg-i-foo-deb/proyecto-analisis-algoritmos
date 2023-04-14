package benchmark

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkNaivOnArray(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		naivOnArray(matricesA[i], matricesB[i])
	}
}

func naivOnArray(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.NAIV_ON_ARRAY, len(matrizA.Datos))()

	multiplicacion_matrices.NaivOnArray(matrizA.Datos, matrizB.Datos)
}
