package cmd

import (
	"generador/pkg/utilidades"

	"github.com/spf13/cobra"
)

var crearNumerosCmd = &cobra.Command{
	Use:   "crear-numeros",
	Short: "Crear números grandes para los algoritmos",
	Long:  "Crea números grandes de diferentes tamaños para ser multiplicados",
	Run:   crearNumeros,
}

func init() {
	rootCmd.AddCommand(crearNumerosCmd)
}

func crearNumeros(cmd *cobra.Command, args []string) {
	utilidades.GenerarNumeros()
}
