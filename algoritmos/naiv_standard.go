package algoritmos

func NaivStandard(a [][]int, b [][]int) [][]int {

	resultado := make([][]int, len(b))

	for i := 0; i < len(a[0]); i++ {
		aux := make([]int, len(b[i]))
		for j := 0; j < len(b); j++ {
			sumatoria := 0
			for k := 0; k < len(a); k++ {
				sumatoria += a[i][k] * b[k][j]
			}
			aux = append(aux, sumatoria)
		}
		resultado = append(resultado, aux)
	}

	return resultado
}
