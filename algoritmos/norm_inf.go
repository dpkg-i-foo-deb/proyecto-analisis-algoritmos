package algoritmos

func NormInf(A [][]int) int {
	N := len(A)
	P := len(A[0])
	Norm := 0
	for i := 0; i < N; i++ {
		aux := 0
		for j := 0; j < P; j++ {
			aux += ObtenerValorAbsoluto(A[i][j])
		}
		Norm = ObtenerMaximo(Norm, aux)
	}
	return Norm
}
