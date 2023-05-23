package cmd

import (
	"generador/pkg/utilidades"
	"log"

	"github.com/spf13/cobra"
)

var generarGraficasGrandesCmd = &cobra.Command{
	Use:   "generar-graficas-grandes",
	Short: "Genera el HTML de las gráficas",
	Long:  "Genera un archivo HTML con las gráficas de multiplicación de números grandes",
	Run:   generarGraficasGrandes,
}

func init() {
	rootCmd.AddCommand(generarGraficasGrandesCmd)
}

func generarGraficasGrandes(cmd *cobra.Command, args []string) {
	log.Println("Leyendo el archivo de resultados...")
	resultados := utilidades.LeerResultadosGrandes()

	log.Println("Generando gráficas...")

	utilidades.GenerarGraficasPromedioMultiplicacionGrandes(resultados)
	log.Println("Gráfica generada en graficas/graficaPromediosGrandes.html")
	utilidades.GenerarGraficaCrecienteMultiplicacionGrandes(resultados)
	log.Println("Gráfica generada en graficas/graficaCrecienteGrandes.html")

}
