package benchmark

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkStrassenWinograd(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		strassenWinograd(matricesA[i], matricesB[i])
	}
}

func strassenWinograd(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.CronometrarMultiplicacionMatrices(modelos.STRASSEN_WINOGRAD, len(matrizA.Datos))()

	multiplicacion_matrices.StrassenWinograd(matrizA.Datos, matrizB.Datos)
}
