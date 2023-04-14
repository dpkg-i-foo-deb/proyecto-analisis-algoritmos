package multiplicacion_matrices

func NaivLoopUnrollingTwo(a [][]int, b [][]int) [][]int {
	resultado := make([][]int, 0)

	if len(b)%2 == 0 {
		for i := 0; i < len(a); i++ {
			filaAux := make([]int, 0)
			for j := 0; j < len(b[i]); j++ {
				suma := 0
				for k := 0; k < len(b); k += 2 {
					suma += (a[i][k] * b[k][j]) + (a[i][k+1] * b[k+1][j])
				}
				filaAux = append(filaAux, suma)
			}
			resultado = append(resultado, filaAux)
		}
	} else {
		ultimoIndice := len(b) - 1
		for i := 0; i < len(a); i++ {
			filaAux := make([]int, 0)
			for j := 0; j < len(b[i]); j++ {
				suma := 0
				for k := 0; k < ultimoIndice; k += 2 {
					suma += (a[i][k] * b[k][j]) + (a[i][k+1] * b[k+1][j])
				}
				filaAux = append(filaAux, suma+(a[i][ultimoIndice]*b[ultimoIndice][j]))
			}
			resultado = append(resultado, filaAux)
		}
	}

	return resultado
}
