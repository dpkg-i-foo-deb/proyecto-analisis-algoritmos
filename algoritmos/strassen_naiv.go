package algoritmos

func StrassenNaiv(a [][]int, b [][]int) [][]int {

	n := len(a)

	if n <= 16 { //El caso base es una matriz de 16x16
		return NaivStandard(a, b)
	}

	mitad := n / 2

	a11 := make([][]int, mitad) // Superior izquierda de a
	a12 := make([][]int, mitad) // Superior derecha de a
	a21 := make([][]int, mitad) // Inferior izquierda de a
	a22 := make([][]int, mitad) // Inferior derecha de a

	b11 := make([][]int, mitad) // Superior izquierda de b
	b12 := make([][]int, mitad) // Superior derecha de b
	b21 := make([][]int, mitad) // Inferior izquierda de b
	b22 := make([][]int, mitad) // Inferior derecha de b

	var i int

	for i = 0; i < mitad; i++ { //Mochar cada matriz en 4 submatrices
		a11[i] = a[i][:mitad]
		a12[i] = a[i][mitad:]
		a21[i] = a[mitad+i][:mitad]
		a22[i] = a[mitad+i][mitad:]

		b11[i] = b[i][:mitad]
		b12[i] = b[i][mitad:]
		b21[i] = b[mitad+i][:mitad]
		b22[i] = b[mitad+i][mitad:]
	}

	//El caso recursivo es... Toda esta locura

	m1 := StrassenNaiv(SumarMatrices(a11, a22), SumarMatrices(b11, b22))
	m2 := StrassenNaiv(SumarMatrices(a21, a22), b11)
	m3 := StrassenNaiv(a11, RestarMatrices(b12, b22))
	m4 := StrassenNaiv(a22, RestarMatrices(b21, b11))
	m5 := StrassenNaiv(SumarMatrices(a11, a12), b22)
	M6 := StrassenNaiv(RestarMatrices(a21, a11), SumarMatrices(b11, b12))
	m7 := StrassenNaiv(RestarMatrices(a12, a22), SumarMatrices(b21, b22))

	//Y aquí mágicamente se calcula la multiplicación de matrices

	c11 := SumarMatrices(RestarMatrices(SumarMatrices(m1, m4), m5), m7) //Superior izquierda
	c12 := SumarMatrices(m3, m5)                                        //Superior derecha
	c21 := SumarMatrices(m2, m4)                                        //Inferior izquierda
	c22 := SumarMatrices(RestarMatrices(SumarMatrices(m1, m3), m2), M6) //Inferior derecha

	c := make([][]int, n)

	for i := range c {
		c[i] = make([]int, n)
	}

	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			c[i][j] = c11[i][j]
			c[i][j+n/2] = c12[i][j]
			c[i+n/2][j] = c21[i][j]
			c[i+n/2][j+n/2] = c22[i][j]
		}
	}

	return c

}
