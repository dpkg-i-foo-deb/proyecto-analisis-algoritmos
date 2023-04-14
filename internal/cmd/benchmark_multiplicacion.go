package cmd

import (
	"generador/internal/benchmark/bmark_multiplicacion_matrices"
	"generador/pkg/modelos"
	"generador/pkg/resultados"
	"generador/pkg/utilidades"
	"log"

	"github.com/spf13/cobra"
)

var benchmarks []func([]modelos.Matriz, []modelos.Matriz)

var benchmarkCmd = &cobra.Command{
	Use:   "benchmark-multiplicacion",
	Short: "Benchmark de los algoritmos de multiplicación",
	Long:  `Ejecuta los algoritmos de multiplicación de matrices con varios tamaños`,
	Run:   benchmark_multiplicacion,
}

func init() {
	rootCmd.AddCommand(benchmarkCmd)

	benchmarks = append(benchmarks, bmark_multiplicacion_matrices.Bmark_iii_parallel_block)
	benchmarks = append(benchmarks, bmark_multiplicacion_matrices.Bmark_iii_sequential_block)
	benchmarks = append(benchmarks, bmark_multiplicacion_matrices.Bmark_iv_3_sequential_block)
	benchmarks = append(benchmarks, bmark_multiplicacion_matrices.Bmark_iv_4_parallel_block)
	benchmarks = append(benchmarks, bmark_multiplicacion_matrices.BmarkNaivKahan)
	benchmarks = append(benchmarks, bmark_multiplicacion_matrices.BmarkNaivLoopUnrollingFour)
	benchmarks = append(benchmarks, bmark_multiplicacion_matrices.BmarkNaivLoopUnrollingThree)
	benchmarks = append(benchmarks, bmark_multiplicacion_matrices.BmarkNaivLoopUnrollingTwo)
	benchmarks = append(benchmarks, bmark_multiplicacion_matrices.BmarkNaivOnArray)
	benchmarks = append(benchmarks, bmark_multiplicacion_matrices.BmarkNaivStandard)
	benchmarks = append(benchmarks, bmark_multiplicacion_matrices.BmarkStrassenNaiv)
	benchmarks = append(benchmarks, bmark_multiplicacion_matrices.BmarkStrassenWinograd)
	benchmarks = append(benchmarks, bmark_multiplicacion_matrices.BmarkV_3_Sequential_Block)
	benchmarks = append(benchmarks, bmark_multiplicacion_matrices.Bmark_V_4_parallel_block)
	benchmarks = append(benchmarks, bmark_multiplicacion_matrices.BmarkWinogradOriginal)
	benchmarks = append(benchmarks, bmark_multiplicacion_matrices.BmarkWinogradScaled)
}

func benchmark_multiplicacion(cmd *cobra.Command, args []string) {

	log.Println("Buscando y leyendo matrices de archivos...")

	matricesA, matricesB := utilidades.LeerMatrices()

	log.Println("Ejecutando benchmarks... Tomará un rato")

	for i := range benchmarks {

		copiaA := make([]modelos.Matriz, len(matricesA), cap(matricesA))

		copiaB := make([]modelos.Matriz, len(matricesB), cap(matricesB))

		copy(copiaA, matricesA)

		copy(copiaB, matricesB)

		benchmarks[i](copiaA, copiaB)

		resultados.ConsolidarMultiplicacionMatrices()
		resultados.EscribirResultadoMultiplicacionMatricesJSON()
		resultados.EscribirResultadoMultiplicacionMatricesTXT()
	}
}
