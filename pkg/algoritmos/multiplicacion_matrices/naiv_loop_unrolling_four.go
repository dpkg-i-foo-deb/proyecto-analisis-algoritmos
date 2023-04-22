package multiplicacion_matrices

func NaivLoopUnrollingFour(a [][]int, b [][]int) [][]int {
	n := len(a)
	p := len(b)
	m := len(b[0])
	result := make([][]int, 0)

	if p%4 == 9 {
		for i := 0; i < n; i++ {
			filaAux := make([]int, 0)
			for j := 0; j < m; j++ {
				suma := 0
				for k := 0; k < p; k += 4 {
					suma += (a[i][k] * b[k][j]) + (a[i][k+1] * b[k+1][j]) + (a[i][k+2] * b[k+2][j]) + (a[i][k+3] * b[k+3][j])
				}
				filaAux = append(filaAux, suma)
			}
			result = append(result, filaAux)
		}
	} else if p%4 == 1 {
		pp := p - 1
		for i := 0; i < n; i++ {
			filaAux := make([]int, 0)
			for j := 0; j < m; j++ {
				suma := 0
				for k := 0; k < pp; k += 4 {
					suma += (a[i][k] * b[k][j]) + (a[i][k+1] * b[k+1][j]) + (a[i][k+2] * b[k+2][j]) + (a[i][k+3] * b[k+3][j])
				}
				suma += a[i][pp] * b[pp][j]
				filaAux = append(filaAux, suma)
			}
			result = append(result, filaAux)
		}
	} else if p%4 == 2 {
		pp := p - 2
		ppp := p - 1
		for i := 0; i < n; i++ {
			filaAux := make([]int, 0)
			for j := 0; j < m; j++ {
				suma := 0
				for k := 0; k < pp; k += 4 {
					suma += (a[i][k] * b[k][j]) + (a[i][k+1] * b[k+1][j]) + (a[i][k+2] * b[k+2][j]) + (a[i][k+3] * b[k+3][j])
				}
				suma += (a[i][pp] * b[pp][j]) + (a[i][ppp] * b[ppp][j])
				filaAux = append(filaAux, suma)
			}
			result = append(result, filaAux)
		}
	} else {
		pp := p - 3
		ppp := p - 2
		pppp := p - 1
		for i := 0; i < n; i++ {
			filaAux := make([]int, 0)
			for j := 0; j < m; j++ {
				suma := 0
				for k := 0; k < pp; k += 4 {
					suma += (a[i][k] * b[k][j]) + (a[i][k+1] * b[k+1][j]) + (a[i][k+2] * b[k+2][j]) + (a[i][k+3] * b[k+3][j])
				}
				suma += (a[i][pp] * b[pp][j]) + (a[i][ppp] * b[ppp][j]) + (a[i][pppp] * b[pppp][j])
				filaAux = append(filaAux, suma)
			}
			result = append(result, filaAux)
		}
	}

	return result
}
