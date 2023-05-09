package multiplicacion_numeros_grandes

func RussianMultiplication(n1, n2 []int, resultado []int) []int {

	for i := 0; i < len(n1); i++ {
		tempA := n1[i]
		tempB := n2[i]
		partialResult := 0

		for tempA > 0 {
			if tempA%2 != 0 {
				partialResult += tempB
			}
			tempA /= 2
			tempB *= 2
		}

		resultado[i] = partialResult
	}

	return resultado
}
