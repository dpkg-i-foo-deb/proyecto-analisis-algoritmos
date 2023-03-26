package cmd

import (
	"generador/utilidades"

	"github.com/spf13/cobra"
)

var crearMatricesCmd = &cobra.Command{
	Use:   "crear-matrices",
	Short: "Crear matrices para los algoritmos",
	Long: `Crea matrices de diferentes tamaños para ser
	multiplicadas por los diferentes algoritmos implementados`,
	Run: crearMatrices,
}

func init() {
	rootCmd.AddCommand(crearMatricesCmd)
}

func crearMatrices(cmd *cobra.Command, args []string) {
	utilidades.GenerarMatrices()
}
