package benchmark

import (
	"generador/pkg/algoritmos/multiplicacion_matrices"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkStrassenNaiv(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		//wg.Add(1)
		strassenNaiv(matricesA[i], matricesB[i])
	}
}

func strassenNaiv(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.STRASSEN_NAIV, len(matrizA.Datos))()
	//defer wg.Done()

	multiplicacion_matrices.StrassenNaiv(matrizA.Datos, matrizB.Datos)
}
