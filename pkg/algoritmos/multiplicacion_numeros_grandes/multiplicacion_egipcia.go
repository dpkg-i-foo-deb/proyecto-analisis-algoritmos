package multiplicacion_numeros_grandes

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
	// Convertir los slices de dígitos en números enteros.
	a := arrayToInt(n1)
	b := arrayToInt(n2)
	result := 0

	for b > 0 {
		if b%2 != 0 {
			result += a
		}
		a = a << 1
		b = b >> 1
	}
	// Convertir el resultado de nuevo en un slice de dígitos y devolverlo.
	return intToSlice(result)
}

/*
Algoritmo de multiplicación egipcia recursiva para números grandes
Entrada: Dos números enteros positivos
Salida: El producto de los dos números
MultiplicacionEgipciaRecursiva es la función que implementa el algoritmo de multiplicación egipcia recursiva
*/

func MultiplicacionEgipciaRecursiva(n1, n2 []int, resultado []int) []int {
	a := arrayToInt(n1)
	b := arrayToInt(n2)
	result := 0
	return intToSlice(multiplicarRecursivamente(a, b, result))
}

func multiplicarRecursivamente(a int, b int, result int) int {
	if b == 0 {
		return result
	}

	if b%2 != 0 {
		result += a
	}

	a = a << 1
	b = b >> 1

	return multiplicarRecursivamente(a, b, result)
}
