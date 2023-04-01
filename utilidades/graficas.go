package utilidades

import (
	"fmt"
	"generador/modelos"
	"log"
	"os"
	"strings"
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

func verificarDirectorioGraficas() {
	if os.MkdirAll("graficas", os.ModePerm) != nil {
		log.Fatal("Error al crear el directorio de salida de gráficas")
	}
}
