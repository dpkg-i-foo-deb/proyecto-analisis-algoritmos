package utilidades

import (
	"fmt"
	"generador/pkg/modelos"
	"log"
	"strconv"
)

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
