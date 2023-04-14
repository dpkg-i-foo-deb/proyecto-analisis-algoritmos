package cmd

import (
	"generador/pkg/utilidades"
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

	log.Println("Generando gráficas...")

	utilidades.GenerarGraficasPromedio(resultados)
	log.Println("Gráfica generada en graficas/graficaPromedios.html")

	utilidades.GenerarGraficasCreciente(resultados)
	log.Println("Gráfica generada en graficas/graficaCreciente.html")

	utilidades.GenerarTabla(resultados)
	log.Println("Tabla generada, datos guardados en graficas/tabla.txt")

	utilidades.GenerarGraficaPunto(resultados)
	log.Println("Gráfica generada en graficas/graficaPunto.html")
}
