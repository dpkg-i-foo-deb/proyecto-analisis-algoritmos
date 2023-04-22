package multiplicacion_matrices

func RestarMatrices(a [][]int, b [][]int) [][]int {

	resultado := make([][]int, len(a))

	for i := range resultado {
		resultado[i] = make([]int, len(a))
	}

	for i := range a {
		for j := range a[i] {
			resultado[i][j] = a[i][j] - b[i][j]
		}
	}

	return resultado
}
