package multiplicacion_numeros_grandes

import (
	"generador/pkg/utilidades"
)

func MultiplicacionDivideVenceras1(num1 []int, num2 []int, resultado []int) []int {

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

	//Mitades

	mitadNum1 := len(num1) / 2
	mitadNum2 := len(num2) / 2

	num1Izda := num1[:mitadNum1]
	num1Dcha := num1[mitadNum1:]

	num2Izda := num2[:mitadNum2]
	num2Dcha := num2[mitadNum2:]

	//Auxiliares
	res1 := make([]int, len(num1Izda)+len(num2Izda))
	res2 := make([]int, len(num1Izda)+len(num2Dcha))
	res3 := make([]int, len(num1Dcha)+len(num2Izda))
	res4 := make([]int, len(num1Dcha)+len(num2Dcha))

	// Se multiplica la mitad izquierda del multiplicando por la mitad derecha del multiplicador

	res1 = MultiplicacionAmericanaIterativa(num1Izda, num2Izda, res1)

	//Se multiplica la mitad izquierda del multiplicando por la mitad derecha del multiplicador

	res2 = MultiplicacionAmericanaIterativa(num1Izda, num2Izda, res2)

	//Se multiplica la mitad derecha del multiplicando por la mitad izquierda del multiplicador

	res3 = MultiplicacionAmericanaIterativa(num1Dcha, num2Izda, res3)

	//Se multiplica la mitad derecha del multiplicando por la mitad derecha del multiplicador

	res4 = MultiplicacionAmericanaIterativa(num1Dcha, num2Dcha, res4)

	res1 = utilidades.RemoverCerosSlice(res1)
	res2 = utilidades.RemoverCerosSlice(res2)
	res3 = utilidades.RemoverCerosSlice(res3)
	res4 = utilidades.RemoverCerosSlice(res4)

	//Desplazar a la izquierda el primer resultado tantas cifras tenga el multiplicador

	for i := 0; i < len(num2); i++ {
		res1 = append(res1, 0)
	}

	//Desplazar a la izquierda el segundo y tercer resultado la mitas de cifras que tenga el multiplicador

	for i := 0; i < len(num2)/2; i++ {
		res2 = append(res2, 0)
		res2 = append(res2, 0)
	}

	//No hay desplazamiento en el cuarto resultado

	//Sumamos los valores

	sumar(resultado, res1)
	sumar(resultado, res2)
	sumar(resultado, res3)
	sumar(resultado, res4)

	return resultado

}

func sumar(arr1, arr2 []int) {
	for i, j := len(arr1), len(arr2); i > 0 && j >= 0; {

		arr1[i] += arr2[i]

		if arr1[i] > 9 {
			arr1[i-1] += arr1[i] / 10
			arr1[i] %= 10
		}

		i--
		j--
	}
}
