package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/utilidades"
	"sync"
)

func BmarkNaivKahan(matricesA []modelos.Matriz, matricesB []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matricesA {
		wg.Add(1)
		go naivKahan(matricesA[i], matricesB[i], wg)
	}
}

func naivKahan(matrizA modelos.Matriz, matrizB modelos.Matriz, wg *sync.WaitGroup) {
	defer utilidades.MedirTiempo(modelos.NAIV_KAHAN, len(matrizA.Datos))()
	defer wg.Done()

	algoritmos.NaivKahan(matrizA.Datos, matrizB.Datos)
}
