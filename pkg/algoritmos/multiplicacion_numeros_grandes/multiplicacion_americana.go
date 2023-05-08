package multiplicacion_numeros_grandes

import "generador/pkg/modelos"

func MultiplicacionAmericanaIterativa(num1 []int, num2 []int, resultado []int) []int {

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

func MultiplicacionAmericanaIterativaEstructuras(num1, num2, resultado *modelos.ListaSimple) *modelos.ListaSimple {

	for i := num2.GetCantidadNodos() - 1; i >= 0; i-- {

		k := resultado.GetCantidadNodos() - (num2.GetCantidadNodos() - i)

		for j := num1.GetCantidadNodos() - 1; j >= 0; j-- {

			resultado.GetNodo(k).Valor += num1.GetValor(j) * num2.GetValor(i)

			if resultado.GetNodo(k).Valor > 9 {
				resultado.GetNodo(k - 1).Valor += resultado.GetNodo(k).Valor / 10
				resultado.GetNodo(k).Valor %= 10
			}

			k--
		}
	}

	return resultado
}
