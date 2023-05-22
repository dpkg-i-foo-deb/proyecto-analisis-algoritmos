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
	long := Max(len(num1), len(num2))
	num1, num2 = IgualarLongitud(num1, num2, long)

	res := make([]int, 0)

	for i := range num1 {
		res = append(res, num1[i]+num2[i])
	}

	for i := len(res) - 1; i > 0; i-- {
		if res[i] > 9 {
			res[i-1] += res[i] / 10
			res[i] %= 10
		}
	}

	return res
}

func RestarArreglos(num1 []int, num2 []int) []int {
	long := Max(len(num1), len(num2))
	num1, num2 = IgualarLongitud(num1, num2, long)

	res := make([]int, 0)

	carry := 0
	for i := len(num1) - 1; i >= 0; i-- {
		resta := num1[i] - num2[i] - carry

		if resta < 0 {
			carry = 1
			resta += 10
		} else {
			carry = 0
		}

		res = append([]int{resta}, res...)
	}

	// Eliminar los ceros a la izquierda del resultado, si los hay
	inicioCeros := 0
	for inicioCeros < len(num1)-1 && res[inicioCeros] == 0 {
		inicioCeros++
	}

	return res[inicioCeros:]

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
