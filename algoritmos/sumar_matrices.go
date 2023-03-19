package algoritmos

func SumarMatrices(a [][]int, b [][]int) [][]int {

	resultado := make([][]int, 0)

	for i := range a {
		resultado = append(resultado, make([]int, len(a)))
		for j := range a[i] {
			resultado[i][j] = a[i][j] + b[i][j]
		}
	}

	return resultado
}
