package multiplicacion_numeros_grandes

func MultiplicacionHindu(n1 []int, n2 []int) []int {
	a := arrayToInt(n1)
	b := arrayToInt(n2)
	resultado := 0

	for i := 0; i < len(n1); i++ {
		for j := 0; j < len(n2); j++ {
			resultado += n1[i] * n2[j] * pow(10, len(n1)-i-1+len(n2)-j-1)
		}
	}

	return intToSlice(resultado)
}

func pow(base int, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
