package algoritmos

func StrassenNaiv(A [][]int, B [][]int) [][]int {
	n := len(A)

	if n == 1 { // base case: 1x1 matrix
		return [][]int{{A[0][0] * B[0][0]}}
	}

	halfN := n / 2

	A11 := make([][]int, halfN) // top left submatrix of A
	A12 := make([][]int, halfN) // top right submatrix of A
	A21 := make([][]int, halfN) // bottom left submatrix of A
	A22 := make([][]int, halfN) // bottom right submatrix of A

	B11 := make([][]int, halfN) // top left submatrix of B
	B12 := make([][]int, halfN) // top right submatrix of B
	B21 := make([][]int, halfN) // bottom left submatrix of B
	B22 := make([][]int, halfN) // bottom right submatrix of B

	var i int

	for i = 0; i < halfN; i++ { // dividing the matrices into 4 submatrices each
		A11[i] = A[i][:halfN]
		A12[i] = A[i][halfN:]
		A21[i] = A[halfN+i][:halfN]
		A22[i] = A[halfN+i][halfN:]

		B11[i] = B[i][:halfN]
		B12[i] = B[i][halfN:]
		B21[i] = B[halfN+i][:halfN]
		B22[i] = B[halfN+i][halfN:]
	}

	M1 := StrassenNaiv(SumarMatrices(A11, A22), SumarMatrices(B11, B22))
	M2 := StrassenNaiv(SumarMatrices(A21, A22), B11)
	M3 := StrassenNaiv(A11, RestarMatrices(B12, B22))
	M4 := StrassenNaiv(A22, RestarMatrices(B21, B11))
	M5 := StrassenNaiv(SumarMatrices(A11, A12), B22)
	M6 := StrassenNaiv(RestarMatrices(A21, A11), SumarMatrices(B11, B12))
	M7 := StrassenNaiv(RestarMatrices(A12, A22), SumarMatrices(B21, B22))

	C11 := SumarMatrices(RestarMatrices(SumarMatrices(M1, M4), M5), M7) //top left
	C12 := SumarMatrices(M3, M5)                                        //top right
	C21 := SumarMatrices(M2, M4)                                        //bottom left
	C22 := SumarMatrices(RestarMatrices(SumarMatrices(M1, M3), M2), M6) //bottom right

	C := make([][]int, n)

	for i := range C {
		C[i] = make([]int, n)
	}

	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			C[i][j] = C11[i][j]
			C[i][j+n/2] = C12[i][j]
			C[i+n/2][j] = C21[i][j]
			C[i+n/2][j+n/2] = C22[i][j]
		}
	}

	return C

}
