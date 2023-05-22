package multiplicacion_numeros_grandes

import "generador/pkg/utilidades"

/*
 * Integrantes:
 *  Stiven Herrera Sierra.
 *  Mateo Estrada
 *  Laura Suárez.
 *  Daniel Eduardo Puerta.
 */
/*
	Algoritmo de multiplicación egipcia iterativa para números grandes
	Entrada: Dos números enteros positivos
	Salida: El producto de los dos números
	MultiplicacionEgipciaIterativa es la función que implementa el algoritmo de multiplicación egipcia recursiva
*/
func MultiplicacionEgipciaIterativa(n1, n2 []int, resultado []int) []int {

	seguir := true

	for seguir {
		if utilidades.SliceIsOdd(n2) {

			resultado = utilidades.SumarArreglos(resultado, n1)
		}

		temp := make([]int, len(n1)+1)

		temp = MultiplicacionInglesaIterativa(n1, []int{2}, temp)

		n1 = temp

		n2 = DividirEstatico(n2, 2)

		if len(n2) == 1 {
			if n2[0] == 0 {
				seguir = false
			}
		}

	}
	// Convertir el resultado de nuevo en un slice de dígitos y devolverlo.
	return resultado
}

/*
Algoritmo de multiplicación egipcia recursiva para números grandes
Entrada: Dos números enteros positivos
Salida: El producto de los dos números
MultiplicacionEgipciaRecursiva es la función que implementa el algoritmo de multiplicación egipcia recursiva
*/

func MultiplicacionEgipciaRecursiva(n1, n2, resultado []int) []int {
	if len(n2) == 1 {

		if n2[0] == 0 {
			return resultado
		}

	}
	if utilidades.SliceIsOdd(n2) {
		resultado = utilidades.SumarArreglos(resultado, n1)
	}

	temp := make([]int, len(n1)+1)

	temp = MultiplicacionInglesaIterativa(n1, []int{2}, temp)

	n1 = temp

	n2 = DividirEstatico(n2, 2)
	return MultiplicacionEgipciaRecursiva(n1, n2, resultado)
}
