package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/tiempo"
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

	algoritmos.StrassenNaiv(matrizA.Datos, matrizB.Datos)
}
