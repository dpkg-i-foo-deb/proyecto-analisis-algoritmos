package multiplicacion_numeros_grandes

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

func MultiplicacionRusaIterativa(n1, n2 []int, resultados []int) []int {
	// Convertir los slices de dígitos en números enteros.
	a := arrayToInt(n1)
	b := arrayToInt(n2)
	resultado := 0
	for a >= 1 {
		if a%2 != 0 {
			resultado += b
		}
		a /= 2
		b *= 2
	}
	// Convertir el resultado de nuevo en un slice de dígitos y devolverlo.
	return intToSlice(resultado)
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

func MultiplicacionRusaRecursiva(n1, n2 []int, resultados []int) []int {
	// Convertir los slices de dígitos en números enteros.
	a := arrayToInt(n1)
	b := arrayToInt(n2)
	// Calcular el resultado utilizando el algoritmo de multiplicación rusa recursiva.
	resultado := multiplicacionRusaRecursivaHelper(a, b)
	// Convertir el resultado de nuevo en un slice de dígitos y devolverlo.
	return intToSlice(resultado)
}

// multiplicacionRusaRecursivaHelper es una función auxiliar que implementa el algoritmo de multiplicación rusa recursiva.
func multiplicacionRusaRecursivaHelper(a, b int) int {
	if a < 1 {
		return 0
	}
	// Calcular el resultado parcial para esta iteración.
	partialResult := 0
	if a%2 != 0 {
		partialResult = b
	}
	// Llamada recurs con a dividido por dos y b multiplicado por dos.
	return partialResult + multiplicacionRusaRecursivaHelper(a/2, b*2)
}

// arrayToInt toma un slice de dígitos y lo convierte en un número entero.
func arrayToInt(arr []int) int {
	num := 0
	// Iterar a través de los dígitos del slice y multiplicar por 10 y sumar el valor de cada dígito.
	for _, v := range arr {
		num = num*10 + v
	}
	// Devolver el número entero resultante.
	return num
}

// intToSlice toma un número entero y lo convierte en un slice de dígitos.
func intToSlice(num int) []int {
	var digitos []int
	// Dividir el número por 10 y tomar el resto en cada iteración para obtener los dígitos del número.
	for num > 0 {
		digitos = append(digitos, num%10)
		num /= 10
	}
	// Invertir el orden de los dígitos en el slice para que estén en el orden correcto.
	for i, j := 0, len(digitos)-1; i < j; i, j = i+1, j-1 {
		digitos[i], digitos[j] = digitos[j], digitos[i]
	}
	// Devolver el slice de dígitos resultante.
	return digitos
}
