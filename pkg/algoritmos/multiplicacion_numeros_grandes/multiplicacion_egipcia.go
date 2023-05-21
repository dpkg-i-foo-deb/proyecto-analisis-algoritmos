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

	for utilidades.SliceGreaterOrEqualOne(n2) {
		if utilidades.SliceIsOdd(n2) {
			Sumar(resultado, n1)
		}

		n1 = MultiplicarPorDos(n1)
		n2 = DividirEstatico(n2, 2)
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
	if !utilidades.SliceGreaterOrEqualOne(n2) {
		return resultado
	}
	if utilidades.SliceIsOdd(n2) {
		Sumar(resultado, n1)
	}
	n1 = MultiplicarPorDos(n1)
	n2 = DividirEstatico(n2, 2)
	return MultiplicacionEgipciaRecursiva(n1, n2, resultado)
}
