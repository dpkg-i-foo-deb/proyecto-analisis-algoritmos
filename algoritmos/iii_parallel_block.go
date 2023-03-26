package algoritmos

import "sync"

func III_ParallelBlock (a, b [][]int) [][]int {
	resultado := make([][]int, len(a))
	size := len(a)
	bsize := 2

	for i := range resultado {
		resultado[i] = make([]int, size)
	}

	var wg sync.WaitGroup
	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {
				wg.Add(1)
				go func (i1, j1, k1 int) {
					for i := i1; i < i1 + bsize && i < size; i++ {
						for j := j1; j < j1 + bsize && j < size; j++ {
							for k := k1; k < k1 + bsize && k < size; k++ {
								resultado[i][j] += a[i][k] * b[k][j];
							};
						}
					}
					wg.Done()
				}(i1, j1, k1)
			}
		}
	}

	wg.Wait()

	return resultado
}