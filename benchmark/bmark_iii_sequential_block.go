package benchmark

import (
	"generador/algoritmos"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func Bmark_iii_sequential_block(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		iii_sequential_block(matricesA[i], matricesB[i])
	}
}

func iii_sequential_block(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.III_SEQUENTIAL_BLOCK, len(matrizA.Datos))()

	algoritmos.III_SequentialBlock(matrizA.Datos, matrizB.Datos)
}
