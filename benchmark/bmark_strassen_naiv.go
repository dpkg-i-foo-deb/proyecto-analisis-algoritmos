package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/tiempo"
	"sync"
)

func BmarkStrassenNaiv(matricesA []modelos.Matriz, matricesB []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matricesA {
		//wg.Add(1)
		strassenNaiv(matricesA[i], matricesB[i], wg)
	}
}

func strassenNaiv(matrizA modelos.Matriz, matrizB modelos.Matriz, wg *sync.WaitGroup) {
	defer tiempo.MedirTiempo(modelos.STRASSEN_NAIV, len(matrizA.Datos))()
	//defer wg.Done()

	algoritmos.StrassenNaiv(matrizA.Datos, matrizB.Datos)
}
