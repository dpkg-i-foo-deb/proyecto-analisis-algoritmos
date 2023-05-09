package multiplicacion_numeros_grandes

import "generador/pkg/modelos"

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

/*
 * Multiplicación Inglesa iterativa usando estructuras dinámicas.
 *
 * Parámetros:
 *  num1: Lista simple de enteros con los dígitos del multiplicando.
 *  num2: Lista simple de enteros con los dígitos del multiplicador.
 *  resultado: Lista simple de enteros con los dígitos para el resultado.
 */
func MultiplicacionInglesaIterativaEstructuras(num1, num2, resultado *modelos.ListaSimple) *modelos.ListaSimple {
	for i := 0; i < num2.GetCantidadNodos(); i++ {
		for j := 0; j < num1.GetCantidadNodos(); j++ {
			resultado.GetNodo(i + j + 1).Valor += num1.GetValor(j) * num2.GetValor(i)
		}
	}

	for i := resultado.GetCantidadNodos() - 1; i > 0; i-- {
		resultado.GetNodo(i - 1).Valor += resultado.GetValor(i) / 10
		resultado.GetNodo(i).Valor %= 10
	}

	return resultado
}

/*
 * Multiplicación Inglesa recursiva usando estructuras dinámicas.
 *
 * Parámetros:
 *  multiplicando: Lista simple de enteros con los dígitos del multiplicando.
 *  multiplicador: Lista simple de enteros con los dígitos del multiplicador.
 *  resultado: Lista simple de enteros con los dígitos para el resultado.
 *  iMultiplicador: Índice auxiliar para recorrer la lista del multiplicador. Debe iniciar en 0.
 *  jMultiplicando: Índice auxiliar para recorrer la lista del multiplicando. Debe iniciar en 0.
 */
func MultiplicacionInglesaRecursivaEstructuras(multiplicando, multiplicador, resultado *modelos.ListaSimple, iMultiplicador, jMultiplicando int) *modelos.ListaSimple {
	//Si ya se recorrieron todos los dígitos del multiplicador.
	if iMultiplicador >= multiplicador.GetCantidadNodos() {
		return resultado
	}

	//Si ya se recorrieron todos los dígitos del multiplicando.
	if jMultiplicando >= multiplicando.GetCantidadNodos()-1 {
		resultado.GetNodo(iMultiplicador + jMultiplicando + 1).Valor += multiplicador.GetValor(iMultiplicador) * multiplicando.GetValor(jMultiplicando)

		MultiplicacionInglesaRecursivaEstructuras(multiplicando, multiplicador, resultado, iMultiplicador+1, 0)

		if resultado.GetValor(iMultiplicador+jMultiplicando+1) > 9 {
			resultado.GetNodo(iMultiplicador + jMultiplicando).Valor += resultado.GetValor(iMultiplicador+jMultiplicando+1) / 10
			resultado.GetNodo(iMultiplicador + jMultiplicando + 1).Valor = resultado.GetValor(iMultiplicador+jMultiplicando+1) % 10
		}

		return resultado
	}

	resultado.GetNodo(iMultiplicador + jMultiplicando + 1).Valor += multiplicador.GetValor(iMultiplicador) * multiplicando.GetValor(jMultiplicando)

	MultiplicacionInglesaRecursivaEstructuras(multiplicando, multiplicador, resultado, iMultiplicador, jMultiplicando+1)

	if resultado.GetValor(iMultiplicador+jMultiplicando+1) > 9 {
		resultado.GetNodo(iMultiplicador + jMultiplicando).Valor += resultado.GetValor(iMultiplicador+jMultiplicando+1) / 10
		resultado.GetNodo(iMultiplicador + jMultiplicando + 1).Valor = resultado.GetValor(iMultiplicador+jMultiplicando+1) % 10
	}

	return resultado
}
