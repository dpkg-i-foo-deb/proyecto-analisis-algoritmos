package main

import (
	"generador/algoritmos"
	"generador/utilidades"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	generar()
}

func generar() {
	rand.Seed(time.Now().UnixNano())

	matriz := utilidades.GenerateMatrix(1000, 1000)

	file, _ := os.Create("Matriz_" + strconv.Itoa(1000) + "x" + strconv.Itoa(1000) + ".json")
	utilidades.WriteMatrix(file, matriz)
	defer file.Close()

	file2, _ := os.Open("Matriz_" + strconv.Itoa(1000) + "x" + strconv.Itoa(1000) + ".json")
	matrizB := utilidades.ReadMatrix(file2)
	defer file2.Close()

	algoritmos.NaivLoopUnrollingTwo(matrizB.Datos, matrizB.Datos)
}
