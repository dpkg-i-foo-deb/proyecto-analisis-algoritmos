package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/tiempo"
	"sync"
)

func BmarkV_3_Sequential_Block(matricesA []modelos.Matriz, matricesB []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matricesA {
		wg.Add(1)
		go v_3_Sequential_Block(matricesA[i], matricesB[i], wg)
	}
}

func v_3_Sequential_Block(matrizA modelos.Matriz, matrizB modelos.Matriz, wg *sync.WaitGroup) {
	defer tiempo.MedirTiempo(modelos.V_3_SEQUENTIAL_BLOCK, len(matrizA.Datos))()
	defer wg.Done()

	algoritmos.V_3_SequentialBlock(matrizA.Datos, matrizB.Datos)

}
