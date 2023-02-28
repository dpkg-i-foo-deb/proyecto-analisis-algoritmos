package utilidades

import (
	"fmt"
	"strconv"
)

func ImprimirMatriz(matriz [][]int) {
	for i := range matriz {

		fmt.Println("")
		for j := range matriz[i] {
			fmt.Print(strconv.FormatInt(int64(matriz[i][j]), 10) + " ")
		}
	}
}
