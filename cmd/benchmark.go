package cmd

import (
	"github.com/spf13/cobra"
)

var benchmarkCmd = &cobra.Command{
	Use:   "benchmark",
	Short: "Benchmark de los algoritmos de multiplicación",
	Long:  `Ejecuta los algoritmos de multiplicación de matrices con varios tamaños`,
	Run:   benchmark,
}

func init() {
	rootCmd.AddCommand(benchmarkCmd)
}

func benchmark(cmd *cobra.Command, args []string) {

}
