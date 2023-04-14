package algoritmos

// Este algoritmo utiliza un tamaño de bloque fijo, siempre es 2
func IV_3_SequentialBlock(A [][]int, B [][]int) [][]int {
	size := len(A)
	bsize := 2
	C := make([][]int, size)
	for i := range C {
		C[i] = make([]int, size)
	}
	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {
				for i := i1; i < ObtenerMinimo(i1+bsize, size); i++ {
					for j := j1; j < ObtenerMinimo(j1+bsize, size); j++ {
						for k := k1; k < ObtenerMinimo(k1+bsize, size); k++ {
							C[i][k] += A[i][j] * B[j][k]
						}
					}
				}
			}
		}
	}
	return C
}
