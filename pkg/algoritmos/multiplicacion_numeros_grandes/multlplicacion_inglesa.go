package multiplicacion_numeros_grandes

/*
 * Integrantes:
 *  Stiven Herrera Sierra.
 *  Mateo Estrada
 *  Laura Suárez.
 *  Daniel Eduardo Puerta.
 */

/*
 * Multiplicación Inglesa iterativa usando estructuras estáticas.
 *
 * Parámetros:
 *  num1: Arreglo de enteros con los dígitos del multiplicando.
 *  num2: Arreglo de enteros con los dígitos del multiplicador.
 *  resultado: Arreglo de enteros con los dígitos para el resultado.
 * 		Sus dimensiones deben ser la suma de la longitud de los arreglos de multiplicando y multiplicador.
 */
func MultiplicacionInglesaIterativa(num1, num2, resultado []int) []int {
	for i := 0; i < len(num2); i++ {
		for j := 0; j < len(num1); j++ {
			resultado[i+j+1] += num1[j] * num2[i]
		}
	}

	for i := len(resultado) - 1; i > 0; i-- {
		resultado[i-1] += resultado[i] / 10
		resultado[i] %= 10
	}

	return resultado
}

/*
 * Multiplicación Inglesa recursiva usando estructuras estáticas.
 *
 * Parámetros:
 *  multiplicando: Arreglo de enteros con los dígitos del multiplicando.
 *  multiplicador: Arreglo de enteros con los dígitos del multiplicador.
 *  resultado: Arreglo de enteros con los dígitos para el resultado.
 * 		Sus dimensiones deben ser la suma de la longitud de los arreglos de multiplicando y multiplicador.
 *  iMultiplicador: Índice auxiliar para recorrer el arreglo del multiplicador. Debe iniciar en 0.
 *  jMultiplicando: Índice auxiliar para recorrer el arreglo del multiplicando. Debe iniciar en 0.
 */
func MultiplicacionInglesaRecursiva(multiplicando, multiplicador, resultado []int, iMultiplicador, jMultiplicando int) []int {
	//Si ya se recorrieron todos los dígitos del multiplicador.
	if iMultiplicador >= len(multiplicador) {
		return resultado
	}

	//Si ya se recorrieron todos los dígitos del multiplicando.
	if jMultiplicando >= len(multiplicando)-1 {
		resultado[iMultiplicador+jMultiplicando+1] += multiplicador[iMultiplicador] * multiplicando[jMultiplicando]

		MultiplicacionInglesaRecursiva(multiplicando, multiplicador, resultado, iMultiplicador+1, 0)

		if resultado[iMultiplicador+jMultiplicando+1] > 9 {
			resultado[iMultiplicador+jMultiplicando] += resultado[iMultiplicador+jMultiplicando+1] / 10
			resultado[iMultiplicador+jMultiplicando+1] = resultado[iMultiplicador+jMultiplicando+1] % 10
		}

		return resultado
	}

	resultado[iMultiplicador+jMultiplicando+1] += multiplicador[iMultiplicador] * multiplicando[jMultiplicando]

	MultiplicacionInglesaRecursiva(multiplicando, multiplicador, resultado, iMultiplicador, jMultiplicando+1)

	if resultado[iMultiplicador+jMultiplicando+1] > 9 {
		resultado[iMultiplicador+jMultiplicando] += resultado[iMultiplicador+jMultiplicando+1] / 10
		resultado[iMultiplicador+jMultiplicando+1] = resultado[iMultiplicador+jMultiplicando+1] % 10
	}

	return resultado
}
