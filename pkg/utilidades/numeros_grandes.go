package utilidades

import (
	"fmt"
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

	n = removerCeros(n)

	for _, valor := range n {

		if valor < 0 || valor > 9 {
			log.Fatal("El slice contiene datos inválidos")
		}

		resultado += fmt.Sprintf("%d", valor)
	}

	return resultado
}

func removerCeros(n []int) []int {

	tope := 0

	for i := 0; i < len(n); i++ {
		if n[i] != 0 {
			tope = i
			i = len(n) - 1
		}
	}

	return n[tope:]
}
