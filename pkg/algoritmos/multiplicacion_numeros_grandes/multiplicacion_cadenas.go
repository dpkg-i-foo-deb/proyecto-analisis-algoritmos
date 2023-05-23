package multiplicacion_numeros_grandes

import (
	"fmt"
	"strings"
)

func MultiplicarUsandoCadenas(num1 string, num2 string) string {

	n1 := len(num1)
	n2 := len(num2)
	resultado := make([]int, n1+n2)

	// Invertir las cadenas para facilitar el cálculo
	num1 = invertirCadena(num1)
	num2 = invertirCadena(num2)

	// Realizar la multiplicación dígito a dígito
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			digito1 := int(num1[i] - '0')
			digito2 := int(num2[j] - '0')
			resultado[i+j] += digito1 * digito2
		}
	}

	//Acomodar los números mayores a 9
	acarreo := 0
	for i := 0; i < n1+n2; i++ {
		resultado[i] += acarreo
		acarreo = resultado[i] / 10
		resultado[i] %= 10
	}

	// Eliminar los ceros no significativos al final
	k := n1 + n2 - 1
	for k > 0 && resultado[k] == 0 {
		k--
	}

	// Construir el resultado final invirtiendo los dígitos
	var builder strings.Builder
	for i := k; i >= 0; i-- {
		builder.WriteString(fmt.Sprintf("%d", resultado[i]))
	}

	return builder.String()
}

func invertirCadena(s string) string {
	runas := []rune(s)
	n := len(runas)
	for i := 0; i < n/2; i++ {
		runas[i], runas[n-i-1] = runas[n-i-1], runas[i]
	}
	return string(runas)
}
