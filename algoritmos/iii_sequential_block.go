package algoritmos

import (
	"math"
	"unsafe"
)

func III_SequentialBlock(a, b [][]int) [][]int {
	size := len(a)
	bsize := calcularTamanoBloque(323)
	result := make([][]int, len(a))

	for i := range result {
		result[i] = make([]int, size)
	}

	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {
				for i := i1; i < i1+bsize && i < size; i++ {
					for j := j1; j < j1+bsize && j < size; j++ {
						for k := k1; k < k1+bsize && k < size; k++ {
							result[i][j] += a[i][k] * b[k][j]
						}
					}
				}
			}
		}
	}

	return result
}

/*
 * Size: Tamaño de las matrices NxN.
 * l1: Tamaño de la memoria caché L1.
 * l2: Tamaño de la memoria caché L2.
 */
func calcularTamanoBloque(cacheSize int) int {
	block_size := int(math.Sqrt(float64(cacheSize / (2 * int(unsafe.Sizeof(int(0)))))))

	return block_size
}
