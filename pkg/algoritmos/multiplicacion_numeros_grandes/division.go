package multiplicacion_numeros_grandes

import "math/big"

// Generado con IA
func DividirEstatico(dividendo []int, divisor int) []int {

	if divisor == 0 || (len(dividendo) == 1 && dividendo[0] == 0) {
		return make([]int, 1)
	}

	// Convertir el arreglo de dígitos en un número grande
	numero := new(big.Int)
	for _, digito := range dividendo {
		// Multiplicar el número por 10 y sumar el dígito
		numero.Mul(numero, big.NewInt(10))
		numero.Add(numero, big.NewInt(int64(digito)))
	}

	// Dividir el número por el divisor usando el método Div
	resultado := new(big.Int)
	resultado.Div(numero, big.NewInt(int64(divisor)))

	// Convertir el resultado en un arreglo de dígitos
	arregloResultado := []int{}
	for resultado.Sign() > 0 {
		// Obtener el último dígito usando el método Mod
		digito := new(big.Int)
		digito.Mod(resultado, big.NewInt(10))

		// Agregar el dígito al principio del arreglo
		arregloResultado = append([]int{int(digito.Int64())}, arregloResultado...)

		// Dividir el resultado por 10 para eliminar el último dígito
		resultado.Div(resultado, big.NewInt(10))
	}

	// Retornar el arreglo de dígitos del resultado
	return arregloResultado
}
