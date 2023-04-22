package multiplicacion_matrices

import (
	"math"
)

func WinogradScaled(A [][]int, B [][]int) [][]int {
	N := len(A)
	P := len(A[0])
	M := len(B[0])

	/* Create scaled copies of A and B */
	CopyA := make([][]int, N)
	for i := range CopyA {
		CopyA[i] = make([]int, P)
	}
	CopyB := make([][]int, P)
	for i := range CopyB {
		CopyB[i] = make([]int, M)
	}

	/* Scaling factors */
	a := NormInf(A)
	b := NormInf(B)
	lambda := math.Floor(0.5 + math.Log(float64(b)/float64(a))/math.Log(4))

	/* Scaling */
	MultiplicarEscalar(A, CopyA, int(math.Pow(2.0, float64(lambda))))
	MultiplicarEscalar(B, CopyB, int(math.Pow(2.0, -float64(lambda))))

	/* Using Winograd with scaled matrices */
	Resultado := WinogradOriginal(CopyA, CopyB)

	return Resultado
}
