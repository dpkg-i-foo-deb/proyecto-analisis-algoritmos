package algoritmos

func StrassenWinograd(a [][]int, b [][]int) [][]int {

	// Calculamos las dimensiones de las matrices
	n := len(a)
	m := n / 2

	// Comprobamos si las matrices tienen tamaño 1
	if len(a) <= 16 {
		return NaivStandard(a, b)
	}

	// Dividimos cada matriz en cuatro submatrices de tamaño m x m
	a11 := make([][]int, m)
	a12 := make([][]int, m)
	a21 := make([][]int, m)
	a22 := make([][]int, m)

	b11 := make([][]int, m)
	b12 := make([][]int, m)
	b21 := make([][]int, m)
	b22 := make([][]int, m)

	for i := 0; i < m; i++ {
		a11[i] = a[i][:m]
		a12[i] = a[i][m:]
		a21[i] = a[i+m][:m]
		a22[i] = a[i+m][m:]

		b11[i] = b[i][:m]
		b12[i] = b[i][m:]
		b21[i] = b[i+m][:m]
		b22[i] = b[i+m][m:]
	}

	// Calcular s1 hasta s10
	s1 := RestarMatrices(b12, b22)
	s2 := SumarMatrices(a11, a12)
	s3 := SumarMatrices(a21, a22)
	s4 := RestarMatrices(b21, b11)
	s5 := SumarMatrices(a11, a22)
	s6 := SumarMatrices(b11, b22)
	s7 := RestarMatrices(a12, a22)
	s8 := SumarMatrices(b21, b22)
	s9 := RestarMatrices(a11, a21)
	s10 := SumarMatrices(b11, b12)

	// Calcular p1 hasta p7
	p1 := StrassenWinograd(a11, s1)
	p2 := StrassenWinograd(s2, b22)
	p3 := StrassenWinograd(s3, b11)
	p4 := StrassenWinograd(a22, s4)
	p5 := StrassenWinograd(s5, s6)
	p6 := StrassenWinograd(s7, s8)
	p7 := StrassenWinograd(s9, s10)

	// Calcular las submatrices de la matriz resultante
	c11 := SumarMatrices(RestarMatrices(SumarMatrices(p5, p4), p2), p6)
	c12 := SumarMatrices(p1, p2)
	c21 := SumarMatrices(p3, p4)
	c22 := RestarMatrices(RestarMatrices(SumarMatrices(p5, p1), p3), p7)

	// Combinar las submatrices en una única matriz
	c := make([][]int, n)
	for i := 0; i < m; i++ {
		c[i] = make([]int, n)
		copy(c[i], c11[i])
		copy(c[i][m:], c12[i])
	}
	for i := n / 2; i < n; i++ {
		c[i] = make([]int, n)
		copy(c[i], c21[i-m])
		copy(c[i][m:], c22[i-m])
	}
	return c
}
