package utilidades

import (
	"log"
	"math/rand"
	"os"
	"time"
)

func GenerarMatrices() {

	log.Println("Generando matrices...")
	defer log.Println("Matrices generadas")

	rand.Seed(time.Now().UnixNano())

	matriz512 := generateMatrix(512, 512)
	matriz1024 := generateMatrix(1024, 1024)
	matriz2048 := generateMatrix(2048, 2028)

	archivo512, err := os.Create("matriz-512x512.json")

	VerificarError(err)

	archivo1024, err := os.Create("matriz-1024x1024.json")

	VerificarError(err)

	archivo2048, err := os.Create("matriz-2048x2048.json")

	VerificarError(err)

	defer archivo512.Close()
	defer archivo1024.Close()
	defer archivo2048.Close()

	writeMatrix(archivo512, matriz512)
	writeMatrix(archivo1024, matriz1024)
	writeMatrix(archivo2048, matriz2048)

}
