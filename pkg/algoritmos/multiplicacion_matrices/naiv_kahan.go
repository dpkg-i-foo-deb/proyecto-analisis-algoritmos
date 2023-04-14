package multiplicacion_matrices

func NaivKahan(a [][]int, b [][]int) [][]int {
	resultado := make([][]int, 0)
	t := 0
	for i := 0; i < len(a); i++ {
		resultado = append(resultado, make([]int, len(a)))
		for j := 0; j < len(b[i]); j++ {
			suma := 0
			error := 0
			for k := 0; k < len(b); k++ {
				error += error + a[i][k]*b[k][j]
				t = suma + error
				error = (suma - t) + error
				suma = t
			}
			resultado[i][j] = suma
		}
	}

	return resultado
}
