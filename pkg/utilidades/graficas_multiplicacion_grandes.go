package utilidades

import (
	"generador/pkg/modelos"
	"os"
	"sort"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var algoritmosMultiplicacionGrandesOrdenados = []string{
	string(modelos.AMERICANA_ITERATIVA_ESTATICO),
	string(modelos.AMERICANA_ITERATIVA_DINAMICO),
	string(modelos.AMERICANA_RECURSIVA_ESTATICO),
	string(modelos.AMERICANA_RECURSIVA_DINAMICO),
	string(modelos.INGLESA_ITERATIVA_ESTATICO),
	string(modelos.INGLESA_ITERATIVA_DINAMICO),
	string(modelos.INGLESA_RECURSIVA_ESTATICO),
	string(modelos.INGLESA_RECURSIVA_DINAMICO),
	string(modelos.RUSA_ITERATIVA),
	string(modelos.RUSA_RECURSIVA),
	string(modelos.HINDU_ITERATIVA),
	string(modelos.HINDU_RECURSIVA),
	string(modelos.EGIPCIA_ITERATIVA),
	string(modelos.EGIPCIA_RECURSIVA),
	string(modelos.KARATSUBA),
	string(modelos.DIVIDE_Y_VENCERAS),
}

func GenerarGraficasPromedioMultiplicacionGrandes(resultados []modelos.ResultadoMultiplicacionNumerosGrandes) {
	verificarDirectorioGraficas()

	promedios := make(map[string]float64)

	titulo := "Promedio de tiempo de ejecución de los algoritmos de multiplicación de números grandes"
	subtitulo := "Se promedian los tiempos de ejecución de cada algoritmo para cada tamaño de número"

	graficaBarras := charts.NewBar()

	graficaBarras.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    titulo,
				Subtitle: subtitulo,
				Right:    "40%",
			}),
	)

	graficaBarras.SetGlobalOptions(opciones...)

	x := []string{}
	series := []opts.BarData{}

	for _, v := range resultados {
		promedios[string(v.Algoritmo)] += float64(v.Duracion)
	}

	for _, aux := range algoritmosMultiplicacionGrandesOrdenados {
		tiempo := promedios[aux] / 8

		x = append(x, aux)
		series = append(series, opts.BarData{Value: tiempo})

	}

	graficaBarras.SetXAxis(x)

	graficaBarras.AddSeries("Tiempo de ejecución", series)

	graficaBarras.XYReversal()

	file, err := os.Create("graficas/graficaPromediosGrandes.html")
	VerificarError(err)

	defer file.Close()

	err = graficaBarras.Render(file)

	VerificarError(err)

}

func GenerarGraficaCrecienteMultiplicacionGrandes(resultados []modelos.ResultadoMultiplicacionNumerosGrandes) {
	verificarDirectorioGraficas()

	var resultados1024 []modelos.ResultadoMultiplicacionNumerosGrandes

	for i := range resultados {
		if resultados[i].N == 1024 {
			resultados1024 = append(resultados1024, resultados[i])
		}

	}

	sort.Slice(resultados1024, OrdenarAscendenteTiempoMultiplicacionGrandes(resultados1024))

	titulo := "Tiempo de ejecución de los algoritmos de multiplicación de números grandes"

	subtitulo := "Tiempos de ejecución para el caso más grande"

	graficaBarras := charts.NewBar()

	graficaBarras.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    titulo,
				Subtitle: subtitulo,
				Right:    "40%",
			}),
	)

	graficaBarras.SetGlobalOptions(opciones...)

	x := []string{}
	series := []opts.BarData{}

	for i := range resultados1024 {
		x = append(x, string(resultados1024[i].Algoritmo))
		series = append(series, opts.BarData{Value: resultados1024[i].Duracion})
	}

	graficaBarras.SetXAxis(x)

	graficaBarras.AddSeries("Tiempo de ejecución", series)

	graficaBarras.XYReversal()

	file, err := os.Create("graficas/graficaCrecienteGrandes.html")
	VerificarError(err)

	defer file.Close()

	err = graficaBarras.Render(file)

	VerificarError(err)
}

func OrdenarAscendenteTiempoMultiplicacionGrandes(a []modelos.ResultadoMultiplicacionNumerosGrandes) func(int, int) bool {
	return func(i, j int) bool {
		return a[i].Duracion > a[j].Duracion
	}
}
