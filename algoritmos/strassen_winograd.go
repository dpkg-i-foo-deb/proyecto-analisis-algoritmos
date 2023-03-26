package algoritmos

import (
	"generador/utilidades"
	"math"
)

func StrassenWinogradOriginal(a, b [][]int) [][]int {
	n := len(a)
	p := len(b)
	m := len(b[0])

	maxSize := utilidades.Max(n, p)
	maxSize = utilidades.Max(maxSize, m)

	if maxSize < 16 {
		maxSize = 16
	}

	k := math.Floor(math.Log(float64(maxSize))/math.Log(2)) - 4

	newM := int(float64(maxSize)*math.Pow(2, -k)) + 1

	newSize := int(float64(newM) * math.Pow(2, k))

	newA := make([][]int, 0)
	newB := make([][]int, 0)

	for i := 0; i < n; i++ {
		filaAux := make([]int, 0)
		for j := 0; j < p; j++ {
			filaAux = append(filaAux, a[i][j])
		}
		newA = append(newA, filaAux)
	}

	for i := 0; i < p; i++ {
		filaAux := make([]int, 0)
		for j := 0; j < len(b[i]); j++ {
			filaAux = append(filaAux, b[i][j])
		}
		newB = append(newB, filaAux)
	}

	result := strassenWinogradStep(newA, newB, newSize, newM)

	return result
}

func strassenWinogradStep(a, b [][]int, n, m int) [][]int {
	var newSize int
	result := make([][]int, n)

	if n%2 == 0 && n > m {
		newSize = n / 2

		a11 := make([][]int, 0)
		a12 := make([][]int, 0)
		a21 := make([][]int, 0)
		a22 := make([][]int, 0)
		b11 := make([][]int, 0)
		b12 := make([][]int, 0)
		b21 := make([][]int, 0)
		b22 := make([][]int, 0)

		for i := 0; i < newSize; i++ {
			filaAux := make([]int, 0)
			for j := 0; j < newSize; j++ {
				filaAux = append(filaAux, a[i][j])
			}
			a11 = append(a11, filaAux)
		}

		for i := 0; i < newSize; i++ {
			filaAux := make([]int, 0)
			for j := 0; j < newSize; j++ {
				filaAux = append(filaAux, a[i][newSize+j])
			}
			a12 = append(a12, filaAux)
		}

		for i := 0; i < newSize; i++ {
			filaAux := make([]int, 0)
			for j := 0; j < newSize; j++ {
				filaAux = append(filaAux, a[newSize+i][j])
			}
			a21 = append(a21, filaAux)
		}

		for i := 0; i < newSize; i++ {
			filaAux := make([]int, 0)
			for j := 0; j < newSize; j++ {
				filaAux = append(filaAux, a[newSize+i][newSize+j])
			}
			a22 = append(a22, filaAux)
		}

		for i := 0; i < newSize; i++ {
			filaAux := make([]int, 0)
			for j := 0; j < newSize; j++ {
				filaAux = append(filaAux, b[i][j])
			}
			b11 = append(b11, filaAux)
		}

		for i := 0; i < newSize; i++ {
			filaAux := make([]int, 0)
			for j := 0; j < newSize; j++ {
				filaAux = append(filaAux, b[i][newSize+j])
			}
			b12 = append(b12, filaAux)
		}

		for i := 0; i < newSize; i++ {
			filaAux := make([]int, 0)
			for j := 0; j < newSize; j++ {
				filaAux = append(filaAux, b[newSize+i][j])
			}
			b21 = append(b21, filaAux)
		}

		for i := 0; i < newSize; i++ {
			filaAux := make([]int, 0)
			for j := 0; j < newSize; j++ {
				filaAux = append(filaAux, b[newSize+i][newSize+j])
			}
			b22 = append(b22, filaAux)
		}

		a1 := RestarMatrices(a11, a21)
		a2 := RestarMatrices(a22, a1)
		b1 := RestarMatrices(b22, b12)
		b2 := SumarMatrices(b1, b11)

		aux1 := strassenWinogradStep(a11, b11, newSize, m)
		aux2 := strassenWinogradStep(a12, b21, newSize, m)
		aux3 := strassenWinogradStep(a2, b2, newSize, m)
		helper1 := SumarMatrices(a21, a22)
		helper2 := SumarMatrices(b12, b11)

		aux4 := strassenWinogradStep(helper1, helper2, newSize, m)
		aux5 := strassenWinogradStep(a1, b1, newSize, m)
		helper1 = RestarMatrices(a12, a2)

		aux6 := strassenWinogradStep(helper1, b22, newSize, m)
		helper1 = RestarMatrices(b21, b2)

		aux7 := strassenWinogradStep(a22, helper1, newSize, m)
		aux8 := SumarMatrices(aux1, aux3)
		aux9 := SumarMatrices(aux8, aux4)

		resultPart11 := SumarMatrices(aux1, aux2)
		resultPart12 := SumarMatrices(aux9, aux6)
		helper1 = SumarMatrices(aux8, aux5)
		resultPart21 := SumarMatrices(helper1, aux7)
		resultPart22 := SumarMatrices(aux9, aux5)

		for i := 0; i < newSize; i++ {
			filaAux := make([]int, 0)
			for j := 0; j < newSize; j++ {
				filaAux = append(filaAux, resultPart11[i][j])
			}
			result[i] = filaAux
		}

		for i := 0; i < newSize; i++ {
			filaAux := make([]int, 0)
			for j := 0; j < newSize; j++ {
				filaAux = append(filaAux, resultPart12[i][j])
			}
			result[i] = filaAux
		}

		for i := 0; i < newSize; i++ {
			filaAux := make([]int, 0)
			for j := 0; j < newSize; j++ {
				filaAux = append(filaAux, resultPart21[i][j])
			}
			result[i] = filaAux
		}

		for i := 0; i < newSize; i++ {
			filaAux := make([]int, 0)
			for j := 0; j < newSize; j++ {
				filaAux = append(filaAux, resultPart22[i][j])
			}
			result[i] = filaAux
		}
	} else {
		result = NaivStandard(a, b)
	}

	return result
}
