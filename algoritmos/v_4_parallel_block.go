package algoritmos

import "sync"

func V_4_ParallelBlock(A [][]int, B [][]int) [][]int {
	size := len(A)
	bsize := 32 // default block size
	C := make([][]int, size)
	for i := range C {
		C[i] = make([]int, size)
	}

	var wg sync.WaitGroup
	wg.Add(size / bsize * size / bsize)

	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			wg.Add(1) // add a waitgroup counter
			go func(i1, j1 int) {
				defer wg.Done()
				for k1 := 0; k1 < size; k1 += bsize {
					for i := i1; i < i1+bsize && i < size; i++ {
						for j := j1; j < j1+bsize && j < size; j++ {
							for k := k1; k < k1+bsize && k < size; k++ {
								C[i][j] += A[i][k] * B[k][j]
							}
						}
					}
				}
			}(i1, j1)
		}
	}

	wg.Wait()
	return C
}
