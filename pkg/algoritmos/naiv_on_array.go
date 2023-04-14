package algoritmos

func NaivOnArray(a [][]int, b [][]int) [][]int {
	resultado := make([][]int, 0)

	for i := 0; i < len(a); i++ {
		resultado = append(resultado, make([]int, len(a)))
		for j := 0; j < len(b[i]); j++ {
			for k := 0; k < len(b); k++ {
				resultado[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return resultado

}
