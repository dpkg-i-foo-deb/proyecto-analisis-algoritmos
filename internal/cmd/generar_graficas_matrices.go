package cmd

import (
	"generador/pkg/utilidades"
	"log"

	"github.com/spf13/cobra"
)

var generarGraficasMatricesCmd = &cobra.Command{
	Use:   "generar-graficas-matrices",
	Short: "Genera el HTML de las gráficas",
	Long:  "Genera un archivo HTML con las gráficas para representar los tiempos de ejecución obtenidos en el archivo resultados.json",
	Run:   generarGraficasMatrices,
}

func init() {
	rootCmd.AddCommand(generarGraficasMatricesCmd)
}

func generarGraficasMatrices(cmd *cobra.Command, args []string) {
	log.Println("Leyendo el archivo de resultados...")
	resultados := utilidades.LeerResultadosMatrices()

	log.Println("Generando gráficas...")

	utilidades.GenerarGraficasPromedioMultiplicacionMatrices(resultados)
	log.Println("Gráfica generada en graficas/graficaPromedios.html")

	utilidades.GenerarGraficasCrecienteMultiplicacionMatrices(resultados)
	log.Println("Gráfica generada en graficas/graficaCreciente.html")

	utilidades.GenerarTablaMultiplicacionMatrices(resultados)
	log.Println("Tabla generada, datos guardados en graficas/tabla.txt")

	utilidades.GenerarGraficaPuntoMultiplicacionMatrices(resultados)
	log.Println("Gráfica generada en graficas/graficaPunto.html")
}
