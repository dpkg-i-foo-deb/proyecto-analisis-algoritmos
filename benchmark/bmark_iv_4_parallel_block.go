package benchmark

import (
	"generador/algoritmos"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func Bmark_iv_4_parallel_block(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		iv_4_parallel_block(matricesA[i], matricesB[i])
	}
}

func iv_4_parallel_block(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.IV_4_PARALLEL_BLOCK, len(matrizA.Datos))()
	algoritmos.IV_4_Pararell_Block(matrizA.Datos, matrizB.Datos)
}
