package utilidades

import (
	"log"
	"strconv"
)

func FormatearEnteroASlice(n int64) []int {

	if n < 0 {
		log.Fatal("El número debe ser mayor a cero")
	}

	resultado := make([]int, 0)

	numByte := []byte(strconv.FormatInt(n, 10))

	for _, valor := range numByte {
		resultado = append(resultado, int(valor))
	}

	return resultado
}

func FormatearSliceAEntero(n []int) int64 {

	var resultado int64 = 0

	for _, valor := range n {

		if valor < 0 || valor > 9 {
			log.Fatal("El slice contiene datos inválidos")
		}

		resultado *= 10 + int64(valor)
	}

	return resultado
}
