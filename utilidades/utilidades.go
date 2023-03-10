package utilidades

import (
	"encoding/json"
	"generador/modelos"
	"math/rand"
	"os"
)

func GenerateMatrix(rows int, cols int) modelos.Matriz {
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

func WriteMatrix(file *os.File, m modelos.Matriz) {
	encoder := json.NewEncoder(file)

	encoder.Encode(m)
}

func ReadMatrix(file *os.File) modelos.Matriz {
	decoder := json.NewDecoder(file)

	matriz := modelos.Matriz{}

	decoder.Decode(&matriz)

	return matriz
}
