package algoritmos

func MultiplicarEscalar(A [][]int, Result [][]int, alpha int) {
	N := len(A)
	M := len(A[0])

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			Result[i][j] = 0
			Result[i][j] = alpha * A[i][j]
		}
	}
}
