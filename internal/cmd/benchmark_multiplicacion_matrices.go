package cmd

import (
	"generador/internal/benchmark/bmark_multiplicacion_matrices"
	"generador/pkg/modelos"
	"generador/pkg/resultados"
	"generador/pkg/utilidades"
	"log"

	"github.com/spf13/cobra"
)

var benchmarksMatrices []func([]modelos.Matriz, []modelos.Matriz)

var benchmarkMatricesCmd = &cobra.Command{
	Use:   "benchmark-multiplicacion-matrices",
	Short: "Benchmark de los algoritmos de multiplicación",
	Long:  `Ejecuta los algoritmos de multiplicación de matrices con varios tamaños`,
	Run:   benchmark_multiplicacion_matrices,
}

func init() {
	rootCmd.AddCommand(benchmarkMatricesCmd)

	benchmarksMatrices = append(benchmarksMatrices, bmark_multiplicacion_matrices.Bmark_iii_parallel_block)
	benchmarksMatrices = append(benchmarksMatrices, bmark_multiplicacion_matrices.Bmark_iii_sequential_block)
	benchmarksMatrices = append(benchmarksMatrices, bmark_multiplicacion_matrices.Bmark_iv_3_sequential_block)
	benchmarksMatrices = append(benchmarksMatrices, bmark_multiplicacion_matrices.Bmark_iv_4_parallel_block)
	benchmarksMatrices = append(benchmarksMatrices, bmark_multiplicacion_matrices.BmarkNaivKahan)
	benchmarksMatrices = append(benchmarksMatrices, bmark_multiplicacion_matrices.BmarkNaivLoopUnrollingFour)
	benchmarksMatrices = append(benchmarksMatrices, bmark_multiplicacion_matrices.BmarkNaivLoopUnrollingThree)
	benchmarksMatrices = append(benchmarksMatrices, bmark_multiplicacion_matrices.BmarkNaivLoopUnrollingTwo)
	benchmarksMatrices = append(benchmarksMatrices, bmark_multiplicacion_matrices.BmarkNaivOnArray)
	benchmarksMatrices = append(benchmarksMatrices, bmark_multiplicacion_matrices.BmarkNaivStandard)
	benchmarksMatrices = append(benchmarksMatrices, bmark_multiplicacion_matrices.BmarkStrassenNaiv)
	benchmarksMatrices = append(benchmarksMatrices, bmark_multiplicacion_matrices.BmarkStrassenWinograd)
	benchmarksMatrices = append(benchmarksMatrices, bmark_multiplicacion_matrices.BmarkV_3_Sequential_Block)
	benchmarksMatrices = append(benchmarksMatrices, bmark_multiplicacion_matrices.Bmark_V_4_parallel_block)
	benchmarksMatrices = append(benchmarksMatrices, bmark_multiplicacion_matrices.BmarkWinogradOriginal)
	benchmarksMatrices = append(benchmarksMatrices, bmark_multiplicacion_matrices.BmarkWinogradScaled)
}

func benchmark_multiplicacion_matrices(cmd *cobra.Command, args []string) {

	log.Println("Buscando y leyendo matrices de archivos...")

	matricesA, matricesB := utilidades.LeerMatrices()

	log.Println("Ejecutando benchmarks... Tomará un rato")

	for i := range benchmarksMatrices {

		copiaA := make([]modelos.Matriz, len(matricesA), cap(matricesA))

		copiaB := make([]modelos.Matriz, len(matricesB), cap(matricesB))

		copy(copiaA, matricesA)

		copy(copiaB, matricesB)

		benchmarksMatrices[i](copiaA, copiaB)

		resultados.ConsolidarMultiplicacionMatrices()
		resultados.EscribirResultadoMultiplicacionMatricesJSON()
		resultados.EscribirResultadoMultiplicacionMatricesTXT()
	}
}
