package algoritmos

func NaivLoopUnrollingThree(a [][]int, b [][]int) [][]int {

	resultado := make([][]int, 0)

	if len(b)%3 == 0 {
		for i := 0; i < len(a); i++ {
			resultado = append(resultado, make([]int, len(a)))
			for j := 0; j < len(b[0]); j++ {
				aux := 0
				for k := 0; k < len(b); k += 3 {
					aux += a[i][k]*b[k][j] + a[i][k+1]*b[k+1][j] + a[i][k+2]*b[k+2][j]
				}
				resultado[i][j] = aux
			}
		}
	} else {
		if len(b)%3 == 1 {
			pp := len(b) - 1
			for i := 0; i < len(a); i++ {
				resultado = append(resultado, make([]int, len(a)))
				for j := 0; j < len(b[0]); j++ {
					aux := 0
					for k := 0; k < pp; k += 3 {
						aux += a[i][k]*b[k][j] + a[i][k+1]*b[k+1][j] + a[i][k+2]*b[k+2][j]
					}
					resultado[i][j] = aux + a[i][pp]*b[pp][j]
				}
			}
		} else {
			pp := len(b) - 2
			ppp := len(b) - 1
			for i := 0; i < len(a); i++ {
				resultado = append(resultado, make([]int, len(a)))
				for j := 0; j < len(b[0]); j++ {
					aux := 0
					for k := 0; k < pp; k += 3 {
						aux += a[i][k]*b[k][j] + a[i][k+1]*b[k+1][j] + a[i][k+2]*b[k+2][j]
					}
					resultado[i][j] = aux + a[i][pp]*b[pp][j] + a[i][ppp]*b[ppp][j]
				}
			}
		}
	}

	return resultado
}
