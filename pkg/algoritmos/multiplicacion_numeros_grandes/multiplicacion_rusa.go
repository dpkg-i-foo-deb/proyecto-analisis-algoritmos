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

		temp := make([]int, len(n2)+1)

		temp = MultiplicacionInglesaIterativa(n2, []int{2}, temp)

		n2 = temp

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

	temp := make([]int, len(n2)+1)

	temp = MultiplicacionInglesaIterativa(n2, []int{2}, temp)

	n2 = temp

	return MultiplicacionRusaRecursiva(n1, n2, resultado)
}
