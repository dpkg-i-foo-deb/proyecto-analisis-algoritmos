package multiplicacion_numeros_grandes

// Generado con IA
func DividirEstatico(dividendo []int, divisor int) []int {
	quotient := make([]int, len(dividendo))
	carry := 0

	for i := 0; i < len(dividendo); i++ {
		digit := dividendo[i] + carry*10
		quotient[i] = digit >> 1 // Use bitwise right shift instead of division by 2
		carry = digit & 1        // Use bitwise AND instead of modulus 2
	}

	// Remove leading zeros, if any
	startIndex := 0
	for startIndex < len(quotient)-1 && quotient[startIndex] == 0 {
		startIndex++
	}

	return quotient[startIndex:]
}
