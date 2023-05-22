package multiplicacion_numeros_grandes

import (
	"generador/pkg/utilidades"
)

/*
 * Integrantes:
 *  Stiven Herrera Sierra.
 *  Mateo Estrada
 *  Laura Suárez.
 *  Daniel Eduardo Puerta.
 */

/*
 * Multiplicación Rusa iterativa usando estructuras estáticas.
 *
 * Parámetros:
 *  n1: Arreglo de enteros con los dígitos del multiplicando.
 *  n2: Arreglo de enteros con los dígitos del multiplicador.
 *  resultados: Arreglo de enteros con los dígitos para el resultado.
 * 		MultiplicacionRusaiterativa toma dos números enteros representados como slices de dígitos y devuelve su producto como un slice de dígitos.
 */

func MultiplicacionRusaIterativa(n1, n2 []int, resultado []int) []int {
	for utilidades.SliceGreaterOrEqualOne(n1) {
		if utilidades.SliceIsOdd(n1) {
			Sumar(resultado, n2)
		}
		n1 = DividirEstatico(n1, 2)
		n2 = MultiplicarPorDos(n2)
	}
	return resultado
}

/*
 * Multiplicación Rusa Recursiva usando estructuras estáticas.
 *
 * Parámetros:
 *  n1: Arreglo de enteros con los dígitos del multiplicando.
 *  n2: Arreglo de enteros con los dígitos del multiplicador.
 *  resultados: Arreglo de enteros con los dígitos para el resultado.
 * 		MultiplicacionRusaRecurs toma dos números enteros representados como slices de dígitos y devuelve su producto como un slice de dígitos.
 */

func MultiplicacionRusaRecursiva(n1, n2, resultado []int) []int {
	if !utilidades.SliceGreaterOrEqualOne(n1) {
		return resultado
	}
	if utilidades.SliceIsOdd(n1) {
		Sumar(resultado, n2)
	}
	n1 = DividirEstatico(n1, 2)
	n2 = MultiplicarPorDos(n2)
	return MultiplicacionRusaRecursiva(n1, n2, resultado)
}

func MultiplicarPorDos(n2 []int) []int {
	carry := 0
	for i := len(n2) - 1; i >= 0; i-- {
		n2[i] = n2[i]*2 + carry
		carry = n2[i] / 10
		n2[i] %= 10
	}
	if carry > 0 {
		n2 = append([]int{carry}, n2...)
	}
	return n2
}

func Sumar(arr1, arr2 []int) {
	for i, j := len(arr1), len(arr2); i > 0 || j > 0; {
		if i > 0 && j > 0 {
			arr1[i-1] += arr2[j-1]
		} else if j > 0 {
			arr1 = append([]int{arr2[j-1]}, arr1...)
		}

		if i > 0 && arr1[i-1] > 9 {
			if i > 1 {
				arr1[i-2] += arr1[i-1] / 10
			} else {
				arr1 = append([]int{arr1[i-1] / 10}, arr1...)
			}
			arr1[i-1] %= 10
		}

		i--
		j--
	}
}
