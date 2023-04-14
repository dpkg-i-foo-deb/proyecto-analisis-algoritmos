package utilidades

import (
	"encoding/json"
	"generador/pkg/modelos"
	"math/rand"
	"os"
)

func generateMatrix(rows int, cols int) modelos.Matriz {
	datos := make([][]int, rows)

	for i := 0; i < rows; i++ {
		datos[i] = make([]int, cols)

		for j := 0; j < cols; j++ {
			datos[i][j] = rand.Intn(10000)
		}
	}

	matriz := modelos.Matriz{Datos: datos}

	return matriz
}

func writeMatrix(file *os.File, m modelos.Matriz) {
	encoder := json.NewEncoder(file)

	encoder.Encode(m)
}

func readMatrix(file *os.File) modelos.Matriz {
	decoder := json.NewDecoder(file)

	matriz := modelos.Matriz{}

	decoder.Decode(&matriz)

	return matriz
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func readResult(file *os.File) []modelos.Resultado {
	decoder := json.NewDecoder(file)

	resultado := []modelos.Resultado{}

	decoder.Decode(&resultado)

	return resultado
}
