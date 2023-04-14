package bmark_multiplicacion_matrices

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkNaivStandard(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		naivStandard(matricesA[i], matricesB[i])
	}
}

func naivStandard(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.CronometrarMultiplicacionMatrices(modelos.NAIV_STANDARD, len(matrizA.Datos))()

	multiplicacion_matrices.NaivStandard(matrizA.Datos, matrizB.Datos)
}
