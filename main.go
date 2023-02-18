package main

import (
	"encoding/json"
	"generador/modelos"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	matriz := generateMatrix(1000, 1000)

	file, _ := os.Create("Matriz_" + strconv.Itoa(1000) + "x" + strconv.Itoa(1000) + ".json")
	writeMatrix(file, matriz)
	defer file.Close()

	file2, _ := os.Open("Matriz_" + strconv.Itoa(1000) + "x" + strconv.Itoa(1000) + ".json")
	readMatrix(file2)
	defer file2.Close()
}

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
