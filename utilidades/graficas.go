package utilidades

import (
	"fmt"
	"generador/modelos"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

const INICIO_PROMEDIOS = `
<html>
<head>
<script type="text/javascript" src="https://www.gstatic.com/charts/loader.js"></script>
<script type="text/javascript">
  google.charts.load("current", {packages:['corechart']});
  google.charts.setOnLoadCallback(drawChart);
  function drawChart() {
`
const GRAFICA_PROMEDIOS = `
var data = google.visualization.arrayToDataTable([
	["Algoritmo", "Duración", { role: "style" } ],
	{{ datos }}
]);

var view = new google.visualization.DataView(data);
view.setColumns([0, 1,
				 { calc: "stringify",
				   sourceColumn: 1,
				   type: "string",
				   role: "annotation" },
				 2]);

var options = {
  title: "{{ titulo }}",
  height: 800,
  bar: {groupWidth: "95%"},
  legend: { position: "none" },
  chartArea:{ left:300 },
};
var chart = new google.visualization.BarChart(document.getElementById("barchart_values"));
chart.draw(view, options);
`

const FIN_PROMEDIOS = `
  }
  </script>
</head>
<body>
  <div id="barchart_values"></div>
</body>
</html>`

var algoritmosOrdenados = []string{
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

func GenerarGraficasPromedio(resultados []modelos.Resultado) {

	verificarDirectorioGraficas()

	promedios := make(map[string]float64)

	for _, v := range resultados {
		promedios[string(v.Algoritmo)] += float64(v.Duracion)
	}

	var datos string

	for _, aux := range algoritmosOrdenados {
		tiempo := promedios[aux] / 12
		datos += fmt.Sprintf("[\"%s\", %.2f, 'color: #76A7FA'],\n", aux, tiempo)
	}

	grafica := GRAFICA_PROMEDIOS
	grafica = strings.Replace(grafica, "{{ titulo }}", "Tiempo promedio de ejecución", -1)
	grafica = strings.Replace(grafica, "{{ datos }}", datos, -1)

	file, err := os.Create("graficas/graficaPromedios.html")
	VerificarError(err)

	defer file.Close()

	fmt.Fprint(file, INICIO_PROMEDIOS+grafica+FIN_PROMEDIOS)
}

func GenerarGraficasCreciente(resultados []modelos.Resultado) {

	pagina := components.NewPage()

	pagina.SetLayout(components.PageFlexLayout)

	verificarDirectorioGraficas()

	var resultados2 []modelos.Resultado
	var resultados4 []modelos.Resultado
	var resultados8 []modelos.Resultado
	var resultados16 []modelos.Resultado
	var resultados32 []modelos.Resultado
	var resultados64 []modelos.Resultado
	var resultados128 []modelos.Resultado
	var resultados256 []modelos.Resultado
	var resultados512 []modelos.Resultado
	var resultados1024 []modelos.Resultado
	var resultados2048 []modelos.Resultado
	var resultados4096 []modelos.Resultado

	var ptr *[]modelos.Resultado

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

	sort.Slice(resultados2, OrdenarAscendenteTiempo(resultados2))
	sort.Slice(resultados4, OrdenarAscendenteTiempo(resultados4))
	sort.Slice(resultados8, OrdenarAscendenteTiempo(resultados8))
	sort.Slice(resultados16, OrdenarAscendenteTiempo(resultados16))
	sort.Slice(resultados32, OrdenarAscendenteTiempo(resultados32))
	sort.Slice(resultados64, OrdenarAscendenteTiempo(resultados64))
	sort.Slice(resultados128, OrdenarAscendenteTiempo(resultados128))
	sort.Slice(resultados256, OrdenarAscendenteTiempo(resultados256))
	sort.Slice(resultados512, OrdenarAscendenteTiempo(resultados512))
	sort.Slice(resultados1024, OrdenarAscendenteTiempo(resultados1024))
	sort.Slice(resultados2048, OrdenarAscendenteTiempo(resultados2048))
	sort.Slice(resultados4096, OrdenarAscendenteTiempo(resultados4096))

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
			charts.WithXAxisOpts(opts.XAxis{
				AxisLabel: &opts.AxisLabel{
					Rotate:       50,
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
		)

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

func verificarDirectorioGraficas() {
	if os.MkdirAll("graficas", os.ModePerm) != nil {
		log.Fatal("Error al crear el directorio de salida de gráficas")
	}
}

func OrdenarAscendenteCantidad(arreglo []modelos.Resultado) func(int, int) bool {
	return func(i, j int) bool {
		if arreglo[i].N > arreglo[j].N {
			return arreglo[i].N > arreglo[j].N

		}
		return arreglo[i].N < arreglo[j].N
	}
}

func OrdenarAscendenteTiempo(arreglo []modelos.Resultado) func(int, int) bool {
	return func(i, j int) bool {
		return arreglo[i].Duracion > arreglo[j].Duracion
	}
}
