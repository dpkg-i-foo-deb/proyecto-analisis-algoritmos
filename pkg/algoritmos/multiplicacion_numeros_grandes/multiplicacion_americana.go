package multiplicacion_numeros_grandes

func MultiplicacionAmericanaIterativa(num1 []int, num2 []int) []int {
	resultado := make([]int, len(num1)+len(num2)*2)

	for i := len(num2) - 1; i >= 0; i-- {

		k := len(resultado) - (len(num2) - i)

		for j := len(num1) - 1; j >= 0; j-- {

			resultado[k] += num1[j] * num2[i]

			if resultado[k] > 9 {
				resultado[k-1] += resultado[k] / 10
				resultado[k] %= 10
			}

			k--
		}
	}

	return resultado
}
