package multiplicacion_numeros_grandes

import "generador/pkg/utilidades"

// Algoritmo de Toom-Cook
func MultiplicacionDivideVenceras2(num1 []int, num2 []int, resultado []int) []int {

	esN1 := false

	//Se toma el tamaño del arreglo más grande como referencia

	if len(num1) > len(num2) {

		esN1 = true
	}

	//Dado el tamaño más grande, que por cierto debe ser par...
	if esN1 {
		if (len(num1)+1)%2 != 0 {
			num1 = append([]int{0}, num1...)
		}
	} else {
		num2 = append([]int{0}, num2...)
	}

	//Ambos arreglos deben tener la misma dimensión

	if esN1 {
		for i := 0; i < len(num1)-len(num2); i++ {
			num2 = append([]int{0}, num2...)
		}
	} else {
		for i := 0; i < len(num2)-len(num1); i++ {
			num1 = append([]int{0}, num1...)
		}
	}

	//Selecciona y multiplica un digito de cada número

	if len(num1) == 1 && len(num2) == 1 {
		return []int{num1[0] * num2[0]}
	}

	//Se define r como la cantidad de veces que se dividen los números

	r := 3

	//Divide los areeglos en sus partes derecha e izquierda

	a, b := dividirArreglo(num1, r)
	c, d := dividirArreglo(num2, r)

	p0 := make([]int, r*2+1)
	p2 := make([]int, r*2+1)

	MultiplicacionDivideVenceras2(a, b, p0)
	MultiplicacionDivideVenceras2(b, d, p2)

	temp1 := utilidades.SumarArreglos(a, b)
	temp2 := utilidades.SumarArreglos(c, d)

	p1 := make([]int, r*2+1)

	MultiplicacionDivideVenceras2(temp1, temp2, p1)

	p1 = utilidades.RestarArreglos(p1, p0)
	p1 = utilidades.RestarArreglos(p1, p2)

	p0 = desplazarIzquierda(p0, 2*r)
	p1 = desplazarIzquierda(p1, r)

	resultadoRetorno := utilidades.SumarArreglos(utilidades.SumarArreglos(p0, p1), p2)

	return resultadoRetorno

}

// Función para dividir un número en partes más pequeñas
func dividirArreglo(num []int, r int) ([]int, []int) {
	if len(num) <= r {
		return []int{0}, num
	}
	return num[:len(num)-r], num[len(num)-r:]
}

// Función para desplazar un número a la izquierda multiplicando por una potencia de 10
func desplazarIzquierda(num []int, desplazamiento int) []int {
	resultado := make([]int, len(num)+desplazamiento)
	copy(resultado[desplazamiento:], num)
	return resultado
}
