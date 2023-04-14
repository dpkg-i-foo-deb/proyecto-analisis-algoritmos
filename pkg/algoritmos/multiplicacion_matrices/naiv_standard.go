package multiplicacion_matrices

func NaivStandard(a [][]int, b [][]int) [][]int {

	resultado := make([][]int, 0)

	for i := 0; i < len(a); i++ {
		aux := make([]int, 0)
		for j := 0; j < len(b[i]); j++ {
			sumatoria := 0
			for k := 0; k < len(b); k++ {
				sumatoria += a[i][k] * b[k][j]
			}
			aux = append(aux, sumatoria)
		}
		resultado = append(resultado, aux)
	}

	return resultado
}
