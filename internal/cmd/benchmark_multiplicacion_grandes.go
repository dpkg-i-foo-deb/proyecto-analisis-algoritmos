package cmd

import (
	"generador/internal/benchmark/bmark_multiplicacion_grandes"
	"generador/pkg/modelos"
	"generador/pkg/resultados"
	"generador/pkg/utilidades"
	"log"

	"github.com/spf13/cobra"
)

var benchmarkGrandes []func([]modelos.NumeroGrande, []modelos.NumeroGrande)

var benchmarkGrandesCmd = &cobra.Command{
	Use:   "benchmark-multiplicacion-grandes",
	Short: "Benchmark de multiplicación de números grandes",
	Long:  "Ejecuta los algoritmos de multiplicación de números grandes",
	Run:   benchmark_multiplicacion_numeros_grandes,
}

func init() {
	rootCmd.AddCommand(benchmarkGrandesCmd)

	benchmarkGrandes = append(benchmarkGrandes, bmark_multiplicacion_grandes.BmarkAmericanaIterativaEstatica)
	benchmarkGrandes = append(benchmarkGrandes, bmark_multiplicacion_grandes.BmarkAmericanaIterativaDinamica)
	benchmarkGrandes = append(benchmarkGrandes, bmark_multiplicacion_grandes.BmarkAmericanaRecursivoEstatico)
	benchmarkGrandes = append(benchmarkGrandes, bmark_multiplicacion_grandes.BmarkAmericanaRecursivoDinamico)

	benchmarkGrandes = append(benchmarkGrandes, bmark_multiplicacion_grandes.BmarkInglesaIterativaEstatica)
	benchmarkGrandes = append(benchmarkGrandes, bmark_multiplicacion_grandes.BmarkInglesaIterativaDinamica)
	benchmarkGrandes = append(benchmarkGrandes, bmark_multiplicacion_grandes.BmarkInglesaRecursivoEstatico)
	benchmarkGrandes = append(benchmarkGrandes, bmark_multiplicacion_grandes.BmarkInglesaRecursivoDinamico)

	benchmarkGrandes = append(benchmarkGrandes, bmark_multiplicacion_grandes.BmarkRusaIterativa)
	benchmarkGrandes = append(benchmarkGrandes, bmark_multiplicacion_grandes.BmarkRusaRecursiva)

	benchmarkGrandes = append(benchmarkGrandes, bmark_multiplicacion_grandes.BmarkHinduIterativa)
	benchmarkGrandes = append(benchmarkGrandes, bmark_multiplicacion_grandes.BmarkHinduRecursiva)

	benchmarkGrandes = append(benchmarkGrandes, bmark_multiplicacion_grandes.BmarkEgipciaIterativo)
	benchmarkGrandes = append(benchmarkGrandes, bmark_multiplicacion_grandes.BmarkEgipciaRecursivo)

	benchmarkGrandes = append(benchmarkGrandes, bmark_multiplicacion_grandes.BmarkKaratsuba)

	benchmarkGrandes = append(benchmarkGrandes, bmark_multiplicacion_grandes.BmarkCadenas)

	benchmarkGrandes = append(benchmarkGrandes, bmark_multiplicacion_grandes.BmarkDivideVenceras)
}

func benchmark_multiplicacion_numeros_grandes(cmd *cobra.Command, args []string) {

	log.Println("Buscando y leyendo números de archivos...")

	numerosA, numerosB := utilidades.LeerNumeros()

	log.Println("Ejecutando benchmarks... Tomará un rato")

	for i := range benchmarkGrandes {

		num1 := make([]modelos.NumeroGrande, len(numerosA))

		num2 := make([]modelos.NumeroGrande, len(numerosB))

		copy(num1, numerosA)

		copy(num2, numerosB)

		benchmarkGrandes[i](num1, num2)

		resultados.ConsolidarMultiplicacionNumerosGrandes()
		resultados.EscribirResultadoMultiplicacionesGrandesJSON()
		resultados.EscribirResultadoMultiplicacionesGrandesTXT()
	}
}
