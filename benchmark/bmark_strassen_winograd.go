package benchmark

import (
	"generador/pkg/algoritmos"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkStrassenWinograd(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		strassenWinograd(matricesA[i], matricesB[i])
	}
}

func strassenWinograd(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.STRASSEN_WINOGRAD, len(matrizA.Datos))()

	algoritmos.StrassenWinograd(matrizA.Datos, matrizB.Datos)
}
