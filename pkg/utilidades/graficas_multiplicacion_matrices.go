package utilidades

import (
	"fmt"
	"generador/pkg/modelos"
	"log"
	"math"
	"os"
	"sort"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/olekukonko/tablewriter"
)

var algoritmosMultiplicacionMatricesOrdenados = []string{
	string(modelos.NAIV_STANDARD),
	string(modelos.NAIV_ON_ARRAY),
	string(modelos.NAIV_KAHAN),
	string(modelos.NAIV_LOOP_UNROLLING_TWO),
	string(modelos.NAIV_LOOP_UNROLLING_THREE),
	string(modelos.NAIV_LOOP_UNROLLING_FOUR),
	string(modelos.WINOGRAD_ORIGINAL),
	string(modelos.WINOGRAD_SCALED),
	string(modelos.STRASSEN_NAIV),
	string(modelos.STRASSEN_WINOGRAD),
	string(modelos.III_SEQUENTIAL_BLOCK),
	string(modelos.III_PARALLEL_BLOCK),
	string(modelos.IV_3_SEQUENTIAL_BLOCK),
	string(modelos.IV_4_PARALLEL_BLOCK),
	string(modelos.V_3_SEQUENTIAL_BLOCK),
	string(modelos.V_4_PARALLEL_BLOCK),
}

var (
	opciones = []charts.GlobalOpts{
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Rotate:       20,
				Show:         true,
				Margin:       10,
				ShowMinLabel: true,
				ShowMaxLabel: true,
			},
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1200px",
			Height: "700px",
		}),
		charts.WithGridOpts(opts.Grid{
			Left:   "20%",
			Bottom: "15%",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "80%"}),
		charts.WithToolboxOpts(opts.Toolbox{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "80%"}),
	}
)

func GenerarGraficasPromedioMultiplicacionMatrices(resultados []modelos.ResultadoMultiplicacionMatrices) {

	verificarDirectorioGraficas()

	promedios := make(map[string]float64)

	titulo := "Promedio de tiempos de ejecución de los algoritmos"
	subtitulo := "Se promedian los tiempos de ejecución de cada algoritmo para cada tamaño de matriz"

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

	for _, aux := range algoritmosMultiplicacionMatricesOrdenados {
		tiempo := promedios[aux] / 12

		x = append(x, aux)
		series = append(series, opts.BarData{Value: tiempo})

	}

	graficaBarras.SetXAxis(x)

	graficaBarras.AddSeries("Tiempo de ejecución", series)

	graficaBarras.XYReversal()

	file, err := os.Create("graficas/graficaPromedios.html")
	VerificarError(err)

	defer file.Close()

	err = graficaBarras.Render(file)

	VerificarError(err)

}

func GenerarGraficasCrecienteMultiplicacionMatrices(resultados []modelos.ResultadoMultiplicacionMatrices) {

	pagina := components.NewPage()

	pagina.SetLayout(components.PageFlexLayout)

	verificarDirectorioGraficas()

	var resultados2 []modelos.ResultadoMultiplicacionMatrices
	var resultados4 []modelos.ResultadoMultiplicacionMatrices
	var resultados8 []modelos.ResultadoMultiplicacionMatrices
	var resultados16 []modelos.ResultadoMultiplicacionMatrices
	var resultados32 []modelos.ResultadoMultiplicacionMatrices
	var resultados64 []modelos.ResultadoMultiplicacionMatrices
	var resultados128 []modelos.ResultadoMultiplicacionMatrices
	var resultados256 []modelos.ResultadoMultiplicacionMatrices
	var resultados512 []modelos.ResultadoMultiplicacionMatrices
	var resultados1024 []modelos.ResultadoMultiplicacionMatrices
	var resultados2048 []modelos.ResultadoMultiplicacionMatrices
	var resultados4096 []modelos.ResultadoMultiplicacionMatrices

	var ptr *[]modelos.ResultadoMultiplicacionMatrices

	for i := range resultados {
		switch resultados[i].N {
		case 2:
			resultados2 = append(resultados2, resultados[i])
		case 4:
			resultados4 = append(resultados4, resultados[i])
		case 8:
			resultados8 = append(resultados8, resultados[i])
		case 16:
			resultados16 = append(resultados16, resultados[i])
		case 32:
			resultados32 = append(resultados32, resultados[i])
		case 64:
			resultados64 = append(resultados64, resultados[i])
		case 128:
			resultados128 = append(resultados128, resultados[i])
		case 256:
			resultados256 = append(resultados256, resultados[i])
		case 512:
			resultados512 = append(resultados512, resultados[i])
		case 1024:
			resultados1024 = append(resultados1024, resultados[i])
		case 2048:
			resultados2048 = append(resultados2048, resultados[i])
		case 4096:
			resultados4096 = append(resultados4096, resultados[i])
		}
	}

	sort.Slice(resultados2, OrdenarAscendenteTiempoMultiplicacionMatrices(resultados2))
	sort.Slice(resultados4, OrdenarAscendenteTiempoMultiplicacionMatrices(resultados4))
	sort.Slice(resultados8, OrdenarAscendenteTiempoMultiplicacionMatrices(resultados8))
	sort.Slice(resultados16, OrdenarAscendenteTiempoMultiplicacionMatrices(resultados16))
	sort.Slice(resultados32, OrdenarAscendenteTiempoMultiplicacionMatrices(resultados32))
	sort.Slice(resultados64, OrdenarAscendenteTiempoMultiplicacionMatrices(resultados64))
	sort.Slice(resultados128, OrdenarAscendenteTiempoMultiplicacionMatrices(resultados128))
	sort.Slice(resultados256, OrdenarAscendenteTiempoMultiplicacionMatrices(resultados256))
	sort.Slice(resultados512, OrdenarAscendenteTiempoMultiplicacionMatrices(resultados512))
	sort.Slice(resultados1024, OrdenarAscendenteTiempoMultiplicacionMatrices(resultados1024))
	sort.Slice(resultados2048, OrdenarAscendenteTiempoMultiplicacionMatrices(resultados2048))
	sort.Slice(resultados4096, OrdenarAscendenteTiempoMultiplicacionMatrices(resultados4096))

	for i := 2; i <= 4096; i *= 2 {

		titulo := fmt.Sprintf("Tiempo de ejecución para matrices de %d x %d", i, i)
		subtitulo := fmt.Sprintf("Matrices de %d x %d utilizando varios algoritmos de multiplicación", i, i)

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

		switch i {
		case 2:
			ptr = &resultados2
		case 4:
			ptr = &resultados4
		case 8:
			ptr = &resultados8
		case 16:
			ptr = &resultados16
		case 32:
			ptr = &resultados32
		case 64:
			ptr = &resultados64
		case 128:
			ptr = &resultados128
		case 256:
			ptr = &resultados256
		case 512:
			ptr = &resultados512
		case 1024:
			ptr = &resultados1024
		case 2048:
			ptr = &resultados2048
		case 4096:
			ptr = &resultados4096

		}

		x := []string{}
		series := []opts.BarData{}

		for i := range *ptr {
			x = append(x, string((*ptr)[i].Algoritmo))
			series = append(series, opts.BarData{Value: (*ptr)[i].Duracion})
		}

		graficaBarras.SetXAxis(x)

		graficaBarras.AddSeries("Tiempo de ejecución", series)

		graficaBarras.XYReversal()

		pagina.AddCharts(graficaBarras)

	}

	file, err := os.Create("graficas/graficaCreciente.html")
	VerificarError(err)

	pagina.Render(file)

	defer file.Close()
}

func GenerarTablaMultiplicacionMatrices(resultados []modelos.ResultadoMultiplicacionMatrices) {

	verificarDirectorioGraficas()

	file, err := os.Create("graficas/tabla.txt")

	VerificarError(err)

	algoritmos := make(map[string][]float64)

	for _, v := range resultados {
		algoritmo := string(modelos.AlgoritmoMultiplicacionMatrices(v.Algoritmo))
		algoritmos[algoritmo] = append(algoritmos[algoritmo], float64(v.Duracion))
	}

	tabla := tablewriter.NewWriter(file)
	tabla.SetHeader([]string{"Algoritmo", "Media", "Rango", "Desviación Estándar", "Varianza"})
	tabla.SetColWidth(50)

	for k, v := range algoritmos {
		media := calcularMedia(v)
		rango := calcularRango(v)
		desviacionEstandar := calcularDesviacionEstandar(v)
		varianza := calcularVarianza(v)

		tabla.Append([]string{k, fmt.Sprintf("%.2f", media), fmt.Sprintf("%.2f", rango), fmt.Sprintf("%.2f", desviacionEstandar), fmt.Sprintf("%.2f", varianza)})
	}

	tabla.Render()
}

func GenerarGraficaPuntoMultiplicacionMatrices(resultados []modelos.ResultadoMultiplicacionMatrices) {

	verificarDirectorioGraficas()

	resultados4096 := []modelos.ResultadoMultiplicacionMatrices{}

	algoritmos := []string{}

	for _, v := range resultados {
		if v.N == 4096 {
			resultados4096 = append(resultados4096, v)
		}
	}

	for _, v := range resultados4096 {
		algoritmos = append(algoritmos, string(v.Algoritmo))
	}

	graficaPunto := charts.NewEffectScatter()

	graficaPunto.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    "Tiempo de ejecución para matrices de 4096 x 4096",
				Subtitle: "Matrices de 4096 x 4096 utilizando varios algoritmos de multiplicación",
				Right:    "40%",
			}),
	)

	graficaPunto.SetGlobalOptions(opciones...)

	x := make([]opts.EffectScatterData, 0)

	for i := range resultados4096 {
		x = append(x, opts.EffectScatterData{Value: resultados4096[i].Duracion})
	}

	graficaPunto.SetXAxis(algoritmos).
		AddSeries("Tiempo de ejecución", x)

	file, err := os.Create("graficas/graficaPunto.html")

	VerificarError(err)

	graficaPunto.Render(file)

	defer file.Close()
}

func calcularMedia(times []float64) float64 {
	sum := 0.0
	for _, time := range times {
		sum += time
	}
	return sum / float64(len(times))
}

func calcularRango(times []float64) float64 {
	max := times[0]
	min := times[0]
	for _, time := range times {
		if time > max {
			max = time
		}
		if time < min {
			min = time
		}
	}
	return max - min
}

func calcularDesviacionEstandar(times []float64) float64 {
	mean := calcularMedia(times)
	sum := 0.0
	for _, time := range times {
		sum += math.Pow(time-mean, 2)
	}
	variance := sum / float64(len(times)-1)
	return math.Sqrt(variance)
}

func calcularVarianza(times []float64) float64 {
	mean := calcularMedia(times)
	sum := 0.0
	for _, time := range times {
		sum += math.Pow(time-mean, 2)
	}
	return sum / float64(len(times)-1)
}

func verificarDirectorioGraficas() {
	if os.MkdirAll("graficas", os.ModePerm) != nil {
		log.Fatal("Error al crear el directorio de salida de gráficas")
	}
}

func OrdenarAscendenteCantidadMultiplicacionMatrices(arreglo []modelos.ResultadoMultiplicacionMatrices) func(int, int) bool {
	return func(i, j int) bool {
		if arreglo[i].N > arreglo[j].N {
			return arreglo[i].N > arreglo[j].N

		}
		return arreglo[i].N < arreglo[j].N
	}
}

func OrdenarAscendenteCantidadMultiplicacionesNumerosGrandes(arreglo []modelos.ResultadoMultiplicacionNumerosGrandes) func(int, int) bool {
	return func(i, j int) bool {
		if arreglo[i].N > arreglo[j].N {
			return arreglo[i].N > arreglo[j].N
		}
		return arreglo[i].N < arreglo[j].N
	}
}

func OrdenarAscendenteTiempoMultiplicacionMatrices(arreglo []modelos.ResultadoMultiplicacionMatrices) func(int, int) bool {
	return func(i, j int) bool {
		return arreglo[i].Duracion > arreglo[j].Duracion
	}
}
