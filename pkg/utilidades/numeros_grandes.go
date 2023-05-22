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

func IgualarLongitud(num1, num2 []int, n int) ([]int, []int) {
	if n == len(num1) && n == len(num2) {
		return num1, num2
	}

	long1 := n - len(num1)
	long2 := n - len(num2)

	n1 := append(make([]int, long1), num1...)
	n2 := append(make([]int, long2), num2...)

	return n1, n2
}

func SumarArreglos(num1, num2 []int) []int {
	len1 := len(num1)
	len2 := len(num2)

	// Determine the longer length
	long := len1
	if len2 > long {
		long = len2
	}

	// Create the result array with the initial capacity
	res := make([]int, long+1) // Increase capacity by 1 to accommodate carry

	carry := 0
	for i := 0; i < long; i++ {
		// Get the digits from num1 and num2, or 0 if index is out of range
		digit1 := 0
		if i < len1 {
			digit1 = num1[len1-1-i]
		}

		digit2 := 0
		if i < len2 {
			digit2 = num2[len2-1-i]
		}

		sum := digit1 + digit2 + carry
		res[long-i] = sum % 10
		carry = sum / 10
	}

	res[0] = carry // Store the final carry, if any

	// Trim leading zeros
	if res[0] == 0 {
		res = res[1:]
	}

	return res
}

func RestarArreglos(num1 []int, num2 []int) []int {
	long := Max(len(num1), len(num2))
	num1, num2 = IgualarLongitud(num1, num2, long)

	res := make([]int, long)

	carry := 0
	for i := long - 1; i >= 0; i-- {
		resta := num1[i] - num2[i] - carry

		if resta < 0 {
			carry = 1
			resta += 10
		} else {
			carry = 0
		}

		res[i] = resta
	}

	// Trim leading zeros
	startIndex := 0
	for startIndex < len(res)-1 && res[startIndex] == 0 {
		startIndex++
	}

	return res[startIndex:]

}

func SliceIsOdd(n []int) bool {
	if len(n) == 0 {
		return false
	}

	return n[len(n)-1]%2 != 0
}

func leerResultadoGrandes(file *os.File) []modelos.ResultadoMultiplicacionNumerosGrandes {
	decoder := json.NewDecoder(file)

	resultado := []modelos.ResultadoMultiplicacionNumerosGrandes{}

	decoder.Decode(&resultado)

	return resultado
}
