package resultados

import (
	"encoding/json"
	"generador/pkg/modelos"
	"generador/pkg/utilidades"
	"sort"
	"strconv"
)

var Resultados []modelos.ResultadoAlgoritmoMultiplicacion

var resultadosIII_parallel_block []modelos.ResultadoAlgoritmoMultiplicacion
var resultadosIII_sequential_block []modelos.ResultadoAlgoritmoMultiplicacion
var resultadosIV_3_sequential_block []modelos.ResultadoAlgoritmoMultiplicacion
var resultadosIV_4_parallel_block []modelos.ResultadoAlgoritmoMultiplicacion
var resultadosNaivKakan []modelos.ResultadoAlgoritmoMultiplicacion
var resultadosNaivLoopUnrollingFour []modelos.ResultadoAlgoritmoMultiplicacion
var resultadosNaivLoopUnrollingThree []modelos.ResultadoAlgoritmoMultiplicacion
var resultadosNaivLoopUnrollingTwo []modelos.ResultadoAlgoritmoMultiplicacion
var resultadosNaivOnArray []modelos.ResultadoAlgoritmoMultiplicacion
var resultadosNaivStandard []modelos.ResultadoAlgoritmoMultiplicacion
var resultadosStrassenNaiv []modelos.ResultadoAlgoritmoMultiplicacion
var resultadosStrassenWinograd []modelos.ResultadoAlgoritmoMultiplicacion
var resultadosV_3_Sequential_block []modelos.ResultadoAlgoritmoMultiplicacion
var resultadosV_4_parallel_block []modelos.ResultadoAlgoritmoMultiplicacion
var resultadosWinogradOriginal []modelos.ResultadoAlgoritmoMultiplicacion
var resultadosWinogradScaled []modelos.ResultadoAlgoritmoMultiplicacion

func Consolidar() {

	resultadosIII_parallel_block = []modelos.ResultadoAlgoritmoMultiplicacion{}
	resultadosIII_sequential_block = []modelos.ResultadoAlgoritmoMultiplicacion{}
	resultadosIV_3_sequential_block = []modelos.ResultadoAlgoritmoMultiplicacion{}
	resultadosIV_4_parallel_block = []modelos.ResultadoAlgoritmoMultiplicacion{}
	resultadosNaivKakan = []modelos.ResultadoAlgoritmoMultiplicacion{}
	resultadosNaivLoopUnrollingFour = []modelos.ResultadoAlgoritmoMultiplicacion{}
	resultadosNaivLoopUnrollingThree = []modelos.ResultadoAlgoritmoMultiplicacion{}
	resultadosNaivLoopUnrollingTwo = []modelos.ResultadoAlgoritmoMultiplicacion{}
	resultadosNaivOnArray = []modelos.ResultadoAlgoritmoMultiplicacion{}
	resultadosNaivStandard = []modelos.ResultadoAlgoritmoMultiplicacion{}
	resultadosStrassenNaiv = []modelos.ResultadoAlgoritmoMultiplicacion{}
	resultadosStrassenWinograd = []modelos.ResultadoAlgoritmoMultiplicacion{}
	resultadosV_3_Sequential_block = []modelos.ResultadoAlgoritmoMultiplicacion{}
	resultadosV_4_parallel_block = []modelos.ResultadoAlgoritmoMultiplicacion{}
	resultadosWinogradOriginal = []modelos.ResultadoAlgoritmoMultiplicacion{}
	resultadosWinogradScaled = []modelos.ResultadoAlgoritmoMultiplicacion{}

	filtrar()
	ordenar()
	reconstruir()
}

func filtrar() {
	for i := range Resultados {
		switch Resultados[i].Algoritmo {
		case modelos.III_PARALLEL_BLOCK:
			resultadosIII_parallel_block = append(resultadosIII_parallel_block, Resultados[i])
		case modelos.III_SEQUENTIAL_BLOCK:
			resultadosIII_sequential_block = append(resultadosIII_sequential_block, Resultados[i])
		case modelos.IV_3_SEQUENTIAL_BLOCK:
			resultadosIV_3_sequential_block = append(resultadosIV_3_sequential_block, Resultados[i])
		case modelos.IV_4_PARALLEL_BLOCK:
			resultadosIV_4_parallel_block = append(resultadosIV_4_parallel_block, Resultados[i])
		case modelos.NAIV_KAHAN:
			resultadosNaivKakan = append(resultadosNaivKakan, Resultados[i])
		case modelos.NAIV_LOOP_UNROLLING_FOUR:
			resultadosNaivLoopUnrollingFour = append(resultadosNaivLoopUnrollingFour, Resultados[i])
		case modelos.NAIV_LOOP_UNROLLING_THREE:
			resultadosNaivLoopUnrollingThree = append(resultadosNaivLoopUnrollingThree, Resultados[i])
		case modelos.NAIV_LOOP_UNROLLING_TWO:
			resultadosNaivLoopUnrollingTwo = append(resultadosNaivLoopUnrollingTwo, Resultados[i])
		case modelos.NAIV_ON_ARRAY:
			resultadosNaivOnArray = append(resultadosNaivOnArray, Resultados[i])
		case modelos.NAIV_STANDARD:
			resultadosNaivStandard = append(resultadosNaivStandard, Resultados[i])
		case modelos.STRASSEN_NAIV:
			resultadosStrassenNaiv = append(resultadosStrassenNaiv, Resultados[i])
		case modelos.STRASSEN_WINOGRAD:
			resultadosStrassenWinograd = append(resultadosStrassenWinograd, Resultados[i])
		case modelos.V_3_SEQUENTIAL_BLOCK:
			resultadosV_3_Sequential_block = append(resultadosV_3_Sequential_block, Resultados[i])
		case modelos.V_4_PARALLEL_BLOCK:
			resultadosV_4_parallel_block = append(resultadosV_4_parallel_block, Resultados[i])
		case modelos.WINOGRAD_ORIGINAL:
			resultadosWinogradOriginal = append(resultadosWinogradOriginal, Resultados[i])
		case modelos.WINOGRAD_SCALED:
			resultadosWinogradScaled = append(resultadosWinogradScaled, Resultados[i])

		}
	}
}

func ordenar() {
	sort.Slice(Resultados, utilidades.OrdenarAscendenteCantidad(Resultados))

	sort.Slice(resultadosIII_parallel_block, utilidades.OrdenarAscendenteCantidad(resultadosIII_parallel_block))
	sort.Slice(resultadosIII_sequential_block, utilidades.OrdenarAscendenteCantidad(resultadosIII_sequential_block))
	sort.Slice(resultadosIV_3_sequential_block, utilidades.OrdenarAscendenteCantidad(resultadosIV_3_sequential_block))
	sort.Slice(resultadosIV_4_parallel_block, utilidades.OrdenarAscendenteCantidad(resultadosIV_4_parallel_block))
	sort.Slice(resultadosNaivKakan, utilidades.OrdenarAscendenteCantidad(resultadosNaivKakan))
	sort.Slice(resultadosNaivLoopUnrollingFour, utilidades.OrdenarAscendenteCantidad(resultadosNaivLoopUnrollingFour))
	sort.Slice(resultadosNaivLoopUnrollingThree, utilidades.OrdenarAscendenteCantidad(resultadosNaivLoopUnrollingThree))
	sort.Slice(resultadosNaivLoopUnrollingTwo, utilidades.OrdenarAscendenteCantidad(resultadosNaivLoopUnrollingTwo))
	sort.Slice(resultadosNaivOnArray, utilidades.OrdenarAscendenteCantidad(resultadosNaivOnArray))
	sort.Slice(resultadosNaivStandard, utilidades.OrdenarAscendenteCantidad(resultadosNaivStandard))
	sort.Slice(resultadosStrassenNaiv, utilidades.OrdenarAscendenteCantidad(resultadosStrassenNaiv))
	sort.Slice(resultadosStrassenWinograd, utilidades.OrdenarAscendenteCantidad(resultadosStrassenWinograd))
	sort.Slice(resultadosV_3_Sequential_block, utilidades.OrdenarAscendenteCantidad(resultadosV_3_Sequential_block))
	sort.Slice(resultadosV_4_parallel_block, utilidades.OrdenarAscendenteCantidad(resultadosV_4_parallel_block))
	sort.Slice(resultadosWinogradOriginal, utilidades.OrdenarAscendenteCantidad(resultadosWinogradOriginal))
	sort.Slice(resultadosWinogradScaled, utilidades.OrdenarAscendenteCantidad(resultadosWinogradScaled))

}

func reconstruir() {
	Resultados = []modelos.ResultadoAlgoritmoMultiplicacion{}

	Resultados = append(Resultados, resultadosNaivStandard...)
	Resultados = append(Resultados, resultadosNaivOnArray...)
	Resultados = append(Resultados, resultadosNaivKakan...)
	Resultados = append(Resultados, resultadosNaivLoopUnrollingThree...)
	Resultados = append(Resultados, resultadosNaivLoopUnrollingTwo...)
	Resultados = append(Resultados, resultadosNaivLoopUnrollingFour...)
	Resultados = append(Resultados, resultadosWinogradOriginal...)
	Resultados = append(Resultados, resultadosWinogradScaled...)
	Resultados = append(Resultados, resultadosStrassenNaiv...)
	Resultados = append(Resultados, resultadosStrassenWinograd...)
	Resultados = append(Resultados, resultadosIII_sequential_block...)
	Resultados = append(Resultados, resultadosIII_parallel_block...)
	Resultados = append(Resultados, resultadosIV_3_sequential_block...)
	Resultados = append(Resultados, resultadosIV_4_parallel_block...)
	Resultados = append(Resultados, resultadosV_3_Sequential_block...)
	Resultados = append(Resultados, resultadosV_4_parallel_block...)

}

func EscribirResultadoTXT() {
	cadena := ""

	for i := range Resultados {
		cadena += Resultados[i].Titulo + " " + strconv.FormatInt(Resultados[i].Duracion, 10) + "\n"
	}

	utilidades.EscribirArchivo("resultados.txt", []byte(cadena))
}

func EscribirResultadoJSON() {
	cadena, err := json.Marshal(Resultados)

	utilidades.VerificarError(err)

	utilidades.EscribirArchivo("resultados.json", cadena)
}
