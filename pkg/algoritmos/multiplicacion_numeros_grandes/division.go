package multiplicacion_numeros_grandes

// Generado con IA
func DividirEstatico(dividendo []int, divisor int) []int {

	quotient := make([]int, len(dividendo))
	carry := 0

	for i := 0; i < len(dividendo); i++ {
		digit := dividendo[i] + carry*10
		quotient[i] = digit / 2
		carry = digit % 2
	}

	// Remove leading zeros, if any
	firstNonZero := 0
	for firstNonZero < len(quotient)-1 && quotient[firstNonZero] == 0 {
		firstNonZero++
	}

	return quotient[firstNonZero:]
}
