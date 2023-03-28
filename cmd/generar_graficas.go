package cmd

import (
	"generador/utilidades"
	"log"

	"github.com/spf13/cobra"
)

var generarGraficasCmd = &cobra.Command{
	Use:   "generar-graficas",
	Short: "Genera el HTML de las gráficas",
	Long:  "Genera un archivo HTML con las gráficas para representar los tiempos de ejecución obtenidos en el archivo resultados.json",
	Run:   generarGraficas,
}

func init() {
	rootCmd.AddCommand(generarGraficasCmd)
}

func generarGraficas(cmd *cobra.Command, args []string) {
	log.Println("Leyendo el archivo de resultados...")
	resultados := utilidades.LeerResultados()

	log.Println("Generando gráfica...")
	utilidades.GenerarGraficasPromedio(resultados)
	log.Println("Gráfica generada en graficas/graficaPromedios.html")
}
