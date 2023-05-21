package utilidades

import (
	"encoding/json"
	"fmt"
	"generador/pkg/modelos"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func generarNumeroGrande(cantidad int) modelos.NumeroGrande {
	datos := make([]int, cantidad)

	for i := 0; i < cantidad; i++ {
		datos[i] = rand.Intn(10)
	}

	numero := modelos.NumeroGrande{Datos: datos}

	return numero
}

func escribirNumero(num modelos.NumeroGrande, archivo *os.File) {
	encoder := json.NewEncoder(archivo)

	encoder.Encode(num)
}

func leerNumero(archivo *os.File) modelos.NumeroGrande {
	decoder := json.NewDecoder(archivo)

	numero := modelos.NumeroGrande{}

	decoder.Decode(&numero)

	return numero
}

func FormatearCadenaASlice(n string) []int {

	resultado := make([]int, 0)

	numByte := []byte(n)

	for _, valor := range numByte {

		num, err := strconv.Atoi(string(valor))

		VerificarError(err)

		resultado = append(resultado, num)
	}

	return resultado
}

func FormatearSliceACadena(n []int) string {

	var resultado = ""

	n = RemoverCerosSlice(n)

	for _, valor := range n {

		if valor < 0 || valor > 9 {
			log.Fatal("El slice contiene datos inválidos")
		}

		resultado += fmt.Sprintf("%d", valor)
	}

	return resultado
}

func RemoverCerosSlice(n []int) []int {

	tope := 0

	for i := 0; i < len(n); i++ {
		if n[i] != 0 {
			tope = i
			i = len(n) - 1
		}
	}

	return n[tope:]
}

func RemoverCerosLista(l *modelos.ListaSimple) *modelos.ListaSimple {

	tope := 0

	for i := 0; i < l.GetCantidadNodos(); i++ {
		if l.GetNodo(i).Valor != 0 {
			tope = i
			i = l.GetCantidadNodos() - 1
		}
	}

	for i := 0; i < tope; i++ {
		l.EliminarInicio()
	}

	return l
}

func FormatearCadenaALista(n string) *modelos.ListaSimple {
	l := modelos.Lista(len(n))

	numByte := []byte(n)

	for i, valor := range numByte {

		num, err := strconv.Atoi(string(valor))

		VerificarError(err)

		l.SetPosicion(i, num)
	}

	return l
}

func FormatearListaACadena(l *modelos.ListaSimple) string {

	var resultado = ""

	l = RemoverCerosLista(l)

	for i := 0; i < l.GetCantidadNodos(); i++ {

		if l.GetNodo(i).Valor < 0 || l.GetNodo(i).Valor > 9 {
			log.Fatal("La lista contiene datos inválidos")
		}

		resultado += fmt.Sprintf("%d", l.GetNodo(i).Valor)
	}

	return resultado

}

// generado con AI
// intToSlice toma un número entero y lo convierte en un slice de dígitos.
func IntToSlice(num int) []int {
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

func SliceIsOdd(n []int) bool {
	// Si el slice está vacío, el número representado es 0
	if len(n) == 0 {
		return false
	}
	// Si el último dígito del número es impar, el número es impar
	if n[len(n)-1]%2 != 0 {
		return true
	}
	// Si el último dígito del número es par, el número es par
	return false
}

func SliceGreaterOrEqualOne(n []int) bool {
	// Si el slice está vacío, el número representado es 0
	if len(n) == 0 {
		return false
	}
	// Si el slice tiene al menos un elemento distinto de 0, el número representado es mayor o igual a 1
	for _, digit := range n {
		if digit != 0 {
			return true
		}
	}
	// Si todos los elementos del slice son 0, el número representado es 0
	return false
}
