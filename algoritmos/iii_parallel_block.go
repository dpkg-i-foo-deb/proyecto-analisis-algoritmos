package algoritmos

import (
	"math"
	"sync"
)

// Este algoritmo utiliza un tamaño de bloque de acuerdo a la memoria caché
func III_ParallelBlock(a, b [][]int) [][]int {
	// Get the dimensions of the matrices
	filas := len(a)
	columnas := len(a[0])

	// Declare the resultado matrix as a slice of slices of integers
	var resultado [][]int = make([][]int, filas)

	// Initialize the result matrix with zeros
	for i := 0; i < filas; i++ {
		resultado[i] = make([]int, columnas)
		for j := 0; j < columnas; j++ {
			resultado[i][j] = 0
		}
	}

	// Encontrar el tamaño del bloque, que debe ser un divisor de las filas
	// Y columnas de la matriz, además de ser cercano al tamaño de la memoria caché
	// Del procesador en KB
	bloque := getBlockSize(len(a), len(b), len(b[0]))

	// En caso de que no se encuentre un tamaño de bloque ideal para la matriz,
	// Se asigna el bloque al número divisor de la mitad de la matriz a
	if bloque == -1 {
		bloque = getMiddleDivisor(len(a))
	}

	// Declare a wait group for synchronization
	var wg sync.WaitGroup

	// Perform matrix multiplication by blocks and store the result
	for i := 0; i < filas; i += bloque {
		for j := 0; j < columnas; j += bloque {
			for k := 0; k < columnas; k += bloque {
				// Add one to the wait group counter
				wg.Add(1)
				// Launch a goroutine to multiply the corresponding blocks of the matrices and add them
				go func(i, j, k int) {
					// Defer the done call to notify the wait group that this goroutine is finished
					defer wg.Done()
					for ii := i; ii < i+bloque; ii++ {
						for jj := j; jj < j+bloque; jj++ {
							for kk := k; kk < k+bloque; kk++ {
								resultado[ii][jj] += a[ii][kk] * b[kk][jj]
							}
						}
					}
				}(i, j, k)
			}
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Return the result matrix
	return resultado
}

// A function that returns the size of the L1 cache in bytes
func getL1CacheSize() int {
	// This is a dummy function that returns a fixed value
	// You can use CPUID or other methods to get the actual value
	return 256 * 1024 // 256 KB
}

// A function that returns the cache line size in bytes
func getCacheLineSize() int {
	// This is a dummy function that returns a fixed value
	// You can use CPUID or other methods to get the actual value
	return 64 // 64 bytes
}

// A function that returns the size of a double in bytes
func getDoubleSize() int {
	// This is a constant value in GO
	return 8 // 8 bytes
}

// A function that returns a suitable block size for matrix multiplication by blocks
func getBlockSize(m, n, p int) int {
	// m, n, p are the dimensions of the matrices A (m x n) and B (n x p)
	// The block size should be smaller than or equal to the square root of the L1 cache size divided by 3
	maxBlockSize := int(math.Sqrt(float64(getL1CacheSize()) / 3.0))
	// The block size should match the partition along the columns of A and the rows of B
	minBlockSize := getDoubleSize()
	// The block size should be a multiple of the cache line size
	cacheLineSize := getCacheLineSize()
	// Find the largest block size that satisfies all the criteria
	for blockSize := maxBlockSize; blockSize >= minBlockSize; blockSize -= cacheLineSize {
		if n%blockSize == 0 && p%blockSize == 0 {
			return blockSize
		}
	}
	// If no suitable block size is found, return -1 as an error indicator
	return -1
}
