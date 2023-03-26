package utilidades

import (
	"fmt"
	"generador/modelos"
	"os"
	"strings"
)

const INICIO = `
<html>
<head>
<script type="text/javascript" src="https://www.gstatic.com/charts/loader.js"></script>
<script type="text/javascript">
  google.charts.load("current", {packages:['corechart']});
  google.charts.setOnLoadCallback(drawChart);
  function drawChart() {
`
const GRAFICA = `
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

const FIN = `
  }
  </script>
</head>
<body>
  <div id="barchart_values"></div>
</body>
</html>`

func GenerarGraficasPromedio(resultados []modelos.Resultado) {
	promedios := make(map[string]float64)

	for _, v := range resultados {
		promedios[v.Titulo] += float64(v.Duracion)
	}

	for v := range promedios {
		promedios[v] /= 12
	}

	var datos string

	for v, p := range promedios {
		datos += fmt.Sprintf("[\"%s\", %f, 'color: #76A7FA'],\n", v, p)
	}

	grafica := GRAFICA
	grafica = strings.Replace(grafica, "{{ titulo }}", "Tiempo promedio de ejecución", -1)
	grafica = strings.Replace(grafica, "{{ datos }}", datos, -1)

	file, _ := os.Create("../graficas/graficaPromedios.html")
	defer file.Close()
	
	fmt.Fprint(file, INICIO+grafica+FIN)
}
