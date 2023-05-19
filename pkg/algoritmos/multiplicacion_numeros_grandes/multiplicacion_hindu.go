package multiplicacion_numeros_grandes

/*
 * Integrantes:
 *  Stiven Herrera Sierra.
 *  Mateo Estrada
 *  Laura Suárez.
 *  Daniel Eduardo Puerta.
 */

/*
 * Multiplicación Hindu iterativa usando estructuras estáticas.
 *  Parámetros:
 *  n1: Arreglo de enteros con los dígitos del multiplicando.
 *  n2: Arreglo de enteros con los dígitos del multiplicador.
 *  resultado: Arreglo de enteros con los dígitos para el resultado.
 * 		Sus dimensiones deben ser la suma de la longitud de los arreglos de multiplicando y multiplicador.
 */

func MultiplicacionHinduIterativa(n1, n2 []int, resultado []int) []int {

	//corremos los dígitos de ambos números de derecha a izquierda
	for i := 0; i < len(n1); i++ {
		for j := 0; j < len(n2); j++ {
			// Multiplicamos cada dígito de n1 con cada dígito de n2
			temp := n1[len(n1)-1-i] * n2[len(n2)-1-j]
			// Almacenamos el resultado en una posición del slice resultado
			resultado[len(resultado)-1-i-j] += temp
		}
	}

	// Llevamos los acarreos hacia la izquierda
	for i := len(resultado) - 1; i > 0; i-- {
		if resultado[i] >= 10 {
			// Llevamos el acarreo al dígito anterior y dejamos el resto en el dígito actual
			resultado[i-1] += resultado[i] / 10
			resultado[i] %= 10
		}
	}

	// Eliminamos cualquier cero que pueda haber al principio del slice resultado
	if resultado[0] == 0 {
		resultado = resultado[1:]
	}

	return resultado
}

/*
 * Multiplicación Hindu Recursiva usando estructuras estáticas.
 *  Parámetros:
 *  n1: Arreglo de enteros con los dígitos del multiplicando.
 *  n2: Arreglo de enteros con los dígitos del multiplicador.
 *  resultado: Arreglo de enteros con los dígitos para el resultado.
 * 		Sus dimensiones deben ser la suma de la longitud de los arreglos de multiplicando y multiplicador.
 */

func MultiplicacionHinduRecursiva(n1, n2 []int, resultado []int, i, j int) []int {
	if i < len(n1) {
		if j < len(n2) {
			temp := n1[len(n1)-1-i] * n2[len(n2)-1-j]
			resultado[len(resultado)-1-i-j] += temp
			return MultiplicacionHinduRecursiva(n1, n2, resultado, i, j+1)
		}
		return MultiplicacionHinduRecursiva(n1, n2, resultado, i+1, 0)
	}

	return LlevarAcarreos(resultado, len(resultado)-1)
}

func LlevarAcarreos(resultado []int, i int) []int {
	if i > 0 {
		if resultado[i] >= 10 {
			resultado[i-1] += resultado[i] / 10
			resultado[i] %= 10
		}
		return LlevarAcarreos(resultado, i-1)
	}

	if resultado[0] == 0 {
		resultado = resultado[1:]
	}

	return resultado
}
