package multiplicacion_matrices

import "sync"

// Este algoritmo utiliza un tamaño de bloque fijo, siempre es 2
func IV_4_Pararell_Block(a, b [][]int) [][]int {
	//Toca usar waitgroup para las gorutinas
	var wg sync.WaitGroup
	//El tamaño serán las filas de a
	size := len(a)
	//Creando el resultado
	c := make([][]int, size)
	for i := 0; i < size; i++ {
		c[i] = make([]int, size)
	}
	//El tamaño del bloque es 2, por eso cada for va de 2 en 2
	for i1 := 0; i1 < size; i1 += 2 {
		for j1 := 0; j1 < size; j1 += 2 {
			for k1 := 0; k1 < size; k1 += 2 {
				// Agregando una gorutina
				wg.Add(1)
				// Cada bloque se ejecuta en un hilo
				go func(i1, j1, k1 int) {
					// Y aquí hacemos la magia
					for i := i1; i < i1+2 && i < size; i++ {
						for j := j1; j < j1+2 && j < size; j++ {
							for k := k1; k < k1+2 && k < size; k++ {
								// Esto es básicamente naiv standard con hilos
								c[i][k] += a[i][j] * b[j][k]
							}
						}
					}
					//Aquí finalizan las gorutinas
					wg.Done()
				}(i1, j1, k1) //Y esto le dice a las gorutinas en qué índices deben empezar
			}
		}
	}
	// El hilo principal se detiene a esperar las gorutinas
	wg.Wait()

	return c
}
