package cmd

import (
	"generador/benchmark"
	"generador/modelos"
	"sync"

	"github.com/spf13/cobra"
)

var benchmarks []func([]modelos.Matriz, []modelos.Matriz, *sync.WaitGroup)

var benchmarkCmd = &cobra.Command{
	Use:   "benchmark-multiplicacion",
	Short: "Benchmark de los algoritmos de multiplicación",
	Long:  `Ejecuta los algoritmos de multiplicación de matrices con varios tamaños`,
	Run:   benchmark_multiplicacion,
}

func init() {
	rootCmd.AddCommand(benchmarkCmd)

	benchmarks = append(benchmarks, benchmark.Bmark_iii_parallel_block)
	benchmarks = append(benchmarks, benchmark.Bmark_iii_sequential_block)
	benchmarks = append(benchmarks, benchmark.Bmark_iv_3_sequential_block)
	benchmarks = append(benchmarks, benchmark.Bmark_iv_4_parallel_block)
	benchmarks = append(benchmarks, benchmark.BmarkNaivKahan)
	benchmarks = append(benchmarks, benchmark.BmarkNaivLoopUnrollingFour)
	benchmarks = append(benchmarks, benchmark.BmarkNaivLoopUnrollingThree)
	benchmarks = append(benchmarks, benchmark.BmarkNaivLoopUnrollingTwo)
	benchmarks = append(benchmarks, benchmark.BmarkNaivOnArray)
	benchmarks = append(benchmarks, benchmark.BmarkNaivStandard)
	benchmarks = append(benchmarks, benchmark.BmarkStrassenNaiv)
	benchmarks = append(benchmarks, benchmark.BmarkStrassenWinograd)
	benchmarks = append(benchmarks, benchmark.BmarkV_3_Sequential_Block)
	benchmarks = append(benchmarks, benchmark.Bmark_V_4_parallel_block)
	benchmarks = append(benchmarks, benchmark.BmarkWinogradOriginal)
	benchmarks = append(benchmarks, benchmark.BmarkWinogradScaled)
}

func benchmark_multiplicacion(cmd *cobra.Command, args []string) {

}
