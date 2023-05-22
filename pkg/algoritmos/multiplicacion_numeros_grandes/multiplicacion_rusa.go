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

	seguir := true

	for seguir {
		if utilidades.SliceIsOdd(n1) {
			resultado = utilidades.SumarArreglos(resultado, n2)
		}
		n1 = DividirEstatico(n1, 2)

		tmp := make([]int, (len(n2)+1)*2)

		tmp = MultiplicacionAmericanaIterativa(n2, []int{2}, tmp)

		tmp = utilidades.RemoverCerosSlice(tmp)

		//n2 = MultiplicarPorDos(n2)
		n2 = tmp

		if len(n1) == 1 {
			if n1[0] == 0 {
				seguir = false
			}
		}
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

	if len(n1) == 1 {

		if n1[0] == 0 {
			return resultado
		}

	}
	if utilidades.SliceIsOdd(n1) {
		resultado = utilidades.SumarArreglos(resultado, n2)
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
