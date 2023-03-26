package benchmark

import (
	"generador/algoritmos"
	"generador/modelos"
	"sync"
)

func Bmark_iii_sequential_block(matrices []modelos.Matriz, wg *sync.WaitGroup) {
	for i := range matrices {
		wg.Add(1)
		go iii_sequential_block(matrices[i], wg)
	}
}

func iii_sequential_block(matriz modelos.Matriz, wg *sync.WaitGroup) {
	defer wg.Done()
	algoritmos.III_SequentialBlock(matriz)
}
