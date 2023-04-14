package benchmark

import (
	"generador/algoritmos"
	"generador/pkg/modelos"
	"generador/tiempo"
)

func Bmark_iv_3_sequential_block(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		iv_3_sequential_block(matricesA[i], matricesB[i])
	}
}

func iv_3_sequential_block(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.IV_3_SEQUENTIAL_BLOCK, len(matrizA.Datos))()

	algoritmos.IV_3_SequentialBlock(matrizA.Datos, matrizB.Datos)
}
