package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"generador/tiempo"
)

func BmarkV_3_Sequential_Block(matricesA []modelos.Matriz, matricesB []modelos.Matriz) {
	for i := range matricesA {
		v_3_Sequential_Block(matricesA[i], matricesB[i])
	}
}

func v_3_Sequential_Block(matrizA modelos.Matriz, matrizB modelos.Matriz) {
	defer tiempo.MedirTiempo(modelos.V_3_SEQUENTIAL_BLOCK, len(matrizA.Datos))()

	algoritmos.V_3_SequentialBlock(matrizA.Datos, matrizB.Datos)

}
