package multiplicacion_numeros_grandes

import (
	"generador/pkg/utilidades"
	"strconv"
)

/*
 * Integrantes:
 *  Stiven Herrera Sierra.
 *  Mateo Estrada
 *  Laura Suárez.
 *  Daniel Eduardo Puerta.
 */

/*
 * Multiplicación de Karatsuba iterativa usando estructuras estáticas.
 *
 * Parámetros:
 *  num1: Arreglo de enteros con los dígitos del multiplicando.
 *  num2: Arreglo de enteros con los dígitos del multiplicador.
 *
 * Este método se basa en Divide y vencerás, calculando recursivamente partes más pequeñas de la multiplicación.
 *
 * La fórmula utilizada es:
 * 	P1*10^((floor(n/2) - n) * 2) + [(num1Left + num1Right)*(num2Left + num2Right) - P1 - P2]* 10^(floor(n/2)) + P2
 */
func MultiplicacionKaratsubaRecursiva(num1, num2 []int) []int {
	//Igualar la longitud de los números añadiendo 0s a la izquierda.
	n := utilidades.Max(len(num1), len(num2))

	num1, num2 = utilidades.IgualarLongitud(num1, num2, n)

	if n == 0 {
		return []int{}
	}

	if n == 1 {
		multiplicacion := strconv.Itoa(num1[0] * num2[0])
		return utilidades.FormatearCadenaASlice(multiplicacion)
	}

	primeraMitad := n / 2 //se usa para dividir el número
	segundaMitad := n - primeraMitad //se usa para la multiplicación de P1 por 10^x

	num1Left := num1[0:primeraMitad]
	num1Right := num1[primeraMitad:]

	num2Left := num2[0:primeraMitad]
	num2Right := num2[primeraMitad:]

	//Cálculo de las partes de la multiplicación de Karatsuba.
	parcialP1 := MultiplicacionKaratsubaRecursiva(num1Left, num2Left)
	parcialP3 := MultiplicacionKaratsubaRecursiva(
		utilidades.SumarArreglos(num1Left, num1Right),
		utilidades.SumarArreglos(num2Left, num2Right))

	p1 := append(parcialP1, make([]int, segundaMitad*2)...)
	p2 := MultiplicacionKaratsubaRecursiva(num1Right, num2Right)
	p3 := append(utilidades.RestarArreglos(utilidades.RestarArreglos(parcialP3, parcialP1), p2), make([]int, segundaMitad)...)

	return utilidades.SumarArreglos(utilidades.SumarArreglos(p1, p3), p2)
}
