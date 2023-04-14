package multiplicacion_matrices

func V_3_SequentialBlock(A [][]int, B [][]int) [][]int {
	size := len(A)
	bsize := len(A)
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
							C[k][i] += A[k][j] * B[j][i]
						}
					}
				}
			}
		}
	}
	return C
}
