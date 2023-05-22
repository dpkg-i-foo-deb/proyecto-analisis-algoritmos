package resultados

import (
	"encoding/json"
	"generador/pkg/modelos"
	"generador/pkg/utilidades"
	"sort"
	"strconv"
)

var ResultadosMultiplicacionMatrices []modelos.ResultadoMultiplicacionMatrices

var resultadosIII_parallel_block []modelos.ResultadoMultiplicacionMatrices
var resultadosIII_sequential_block []modelos.ResultadoMultiplicacionMatrices
var resultadosIV_3_sequential_block []modelos.ResultadoMultiplicacionMatrices
var resultadosIV_4_parallel_block []modelos.ResultadoMultiplicacionMatrices
var resultadosNaivKakan []modelos.ResultadoMultiplicacionMatrices
var resultadosNaivLoopUnrollingFour []modelos.ResultadoMultiplicacionMatrices
var resultadosNaivLoopUnrollingThree []modelos.ResultadoMultiplicacionMatrices
var resultadosNaivLoopUnrollingTwo []modelos.ResultadoMultiplicacionMatrices
var resultadosNaivOnArray []modelos.ResultadoMultiplicacionMatrices
var resultadosNaivStandard []modelos.ResultadoMultiplicacionMatrices
var resultadosStrassenNaiv []modelos.ResultadoMultiplicacionMatrices
var resultadosStrassenWinograd []modelos.ResultadoMultiplicacionMatrices
var resultadosV_3_Sequential_block []modelos.ResultadoMultiplicacionMatrices
var resultadosV_4_parallel_block []modelos.ResultadoMultiplicacionMatrices
var resultadosWinogradOriginal []modelos.ResultadoMultiplicacionMatrices
var resultadosWinogradScaled []modelos.ResultadoMultiplicacionMatrices

func ConsolidarMultiplicacionMatrices() {

	resultadosIII_parallel_block = []modelos.ResultadoMultiplicacionMatrices{}
	resultadosIII_sequential_block = []modelos.ResultadoMultiplicacionMatrices{}
	resultadosIV_3_sequential_block = []modelos.ResultadoMultiplicacionMatrices{}
	resultadosIV_4_parallel_block = []modelos.ResultadoMultiplicacionMatrices{}
	resultadosNaivKakan = []modelos.ResultadoMultiplicacionMatrices{}
	resultadosNaivLoopUnrollingFour = []modelos.ResultadoMultiplicacionMatrices{}
	resultadosNaivLoopUnrollingThree = []modelos.ResultadoMultiplicacionMatrices{}
	resultadosNaivLoopUnrollingTwo = []modelos.ResultadoMultiplicacionMatrices{}
	resultadosNaivOnArray = []modelos.ResultadoMultiplicacionMatrices{}
	resultadosNaivStandard = []modelos.ResultadoMultiplicacionMatrices{}
	resultadosStrassenNaiv = []modelos.ResultadoMultiplicacionMatrices{}
	resultadosStrassenWinograd = []modelos.ResultadoMultiplicacionMatrices{}
	resultadosV_3_Sequential_block = []modelos.ResultadoMultiplicacionMatrices{}
	resultadosV_4_parallel_block = []modelos.ResultadoMultiplicacionMatrices{}
	resultadosWinogradOriginal = []modelos.ResultadoMultiplicacionMatrices{}
	resultadosWinogradScaled = []modelos.ResultadoMultiplicacionMatrices{}

	filtrarMultiplicacionMatrices()
	ordenarMultiplicacionMatrices()
	reconstruir()
}

func filtrarMultiplicacionMatrices() {
	for i := range ResultadosMultiplicacionMatrices {
		switch ResultadosMultiplicacionMatrices[i].Algoritmo {
		case modelos.III_PARALLEL_BLOCK:
			resultadosIII_parallel_block = append(resultadosIII_parallel_block, ResultadosMultiplicacionMatrices[i])
		case modelos.III_SEQUENTIAL_BLOCK:
			resultadosIII_sequential_block = append(resultadosIII_sequential_block, ResultadosMultiplicacionMatrices[i])
		case modelos.IV_3_SEQUENTIAL_BLOCK:
			resultadosIV_3_sequential_block = append(resultadosIV_3_sequential_block, ResultadosMultiplicacionMatrices[i])
		case modelos.IV_4_PARALLEL_BLOCK:
			resultadosIV_4_parallel_block = append(resultadosIV_4_parallel_block, ResultadosMultiplicacionMatrices[i])
		case modelos.NAIV_KAHAN:
			resultadosNaivKakan = append(resultadosNaivKakan, ResultadosMultiplicacionMatrices[i])
		case modelos.NAIV_LOOP_UNROLLING_FOUR:
			resultadosNaivLoopUnrollingFour = append(resultadosNaivLoopUnrollingFour, ResultadosMultiplicacionMatrices[i])
		case modelos.NAIV_LOOP_UNROLLING_THREE:
			resultadosNaivLoopUnrollingThree = append(resultadosNaivLoopUnrollingThree, ResultadosMultiplicacionMatrices[i])
		case modelos.NAIV_LOOP_UNROLLING_TWO:
			resultadosNaivLoopUnrollingTwo = append(resultadosNaivLoopUnrollingTwo, ResultadosMultiplicacionMatrices[i])
		case modelos.NAIV_ON_ARRAY:
			resultadosNaivOnArray = append(resultadosNaivOnArray, ResultadosMultiplicacionMatrices[i])
		case modelos.NAIV_STANDARD:
			resultadosNaivStandard = append(resultadosNaivStandard, ResultadosMultiplicacionMatrices[i])
		case modelos.STRASSEN_NAIV:
			resultadosStrassenNaiv = append(resultadosStrassenNaiv, ResultadosMultiplicacionMatrices[i])
		case modelos.STRASSEN_WINOGRAD:
			resultadosStrassenWinograd = append(resultadosStrassenWinograd, ResultadosMultiplicacionMatrices[i])
		case modelos.V_3_SEQUENTIAL_BLOCK:
			resultadosV_3_Sequential_block = append(resultadosV_3_Sequential_block, ResultadosMultiplicacionMatrices[i])
		case modelos.V_4_PARALLEL_BLOCK:
			resultadosV_4_parallel_block = append(resultadosV_4_parallel_block, ResultadosMultiplicacionMatrices[i])
		case modelos.WINOGRAD_ORIGINAL:
			resultadosWinogradOriginal = append(resultadosWinogradOriginal, ResultadosMultiplicacionMatrices[i])
		case modelos.WINOGRAD_SCALED:
			resultadosWinogradScaled = append(resultadosWinogradScaled, ResultadosMultiplicacionMatrices[i])

		}
	}
}

func ordenarMultiplicacionMatrices() {
	sort.Slice(ResultadosMultiplicacionMatrices, utilidades.OrdenarAscendenteCantidadMultiplicacionMatrices(ResultadosMultiplicacionMatrices))

	sort.Slice(resultadosIII_parallel_block, utilidades.OrdenarAscendenteCantidadMultiplicacionMatrices(resultadosIII_parallel_block))
	sort.Slice(resultadosIII_sequential_block, utilidades.OrdenarAscendenteCantidadMultiplicacionMatrices(resultadosIII_sequential_block))
	sort.Slice(resultadosIV_3_sequential_block, utilidades.OrdenarAscendenteCantidadMultiplicacionMatrices(resultadosIV_3_sequential_block))
	sort.Slice(resultadosIV_4_parallel_block, utilidades.OrdenarAscendenteCantidadMultiplicacionMatrices(resultadosIV_4_parallel_block))
	sort.Slice(resultadosNaivKakan, utilidades.OrdenarAscendenteCantidadMultiplicacionMatrices(resultadosNaivKakan))
	sort.Slice(resultadosNaivLoopUnrollingFour, utilidades.OrdenarAscendenteCantidadMultiplicacionMatrices(resultadosNaivLoopUnrollingFour))
	sort.Slice(resultadosNaivLoopUnrollingThree, utilidades.OrdenarAscendenteCantidadMultiplicacionMatrices(resultadosNaivLoopUnrollingThree))
	sort.Slice(resultadosNaivLoopUnrollingTwo, utilidades.OrdenarAscendenteCantidadMultiplicacionMatrices(resultadosNaivLoopUnrollingTwo))
	sort.Slice(resultadosNaivOnArray, utilidades.OrdenarAscendenteCantidadMultiplicacionMatrices(resultadosNaivOnArray))
	sort.Slice(resultadosNaivStandard, utilidades.OrdenarAscendenteCantidadMultiplicacionMatrices(resultadosNaivStandard))
	sort.Slice(resultadosStrassenNaiv, utilidades.OrdenarAscendenteCantidadMultiplicacionMatrices(resultadosStrassenNaiv))
	sort.Slice(resultadosStrassenWinograd, utilidades.OrdenarAscendenteCantidadMultiplicacionMatrices(resultadosStrassenWinograd))
	sort.Slice(resultadosV_3_Sequential_block, utilidades.OrdenarAscendenteCantidadMultiplicacionMatrices(resultadosV_3_Sequential_block))
	sort.Slice(resultadosV_4_parallel_block, utilidades.OrdenarAscendenteCantidadMultiplicacionMatrices(resultadosV_4_parallel_block))
	sort.Slice(resultadosWinogradOriginal, utilidades.OrdenarAscendenteCantidadMultiplicacionMatrices(resultadosWinogradOriginal))
	sort.Slice(resultadosWinogradScaled, utilidades.OrdenarAscendenteCantidadMultiplicacionMatrices(resultadosWinogradScaled))

}

func reconstruir() {
	ResultadosMultiplicacionMatrices = []modelos.ResultadoMultiplicacionMatrices{}

	ResultadosMultiplicacionMatrices = append(ResultadosMultiplicacionMatrices, resultadosNaivStandard...)
	ResultadosMultiplicacionMatrices = append(ResultadosMultiplicacionMatrices, resultadosNaivOnArray...)
	ResultadosMultiplicacionMatrices = append(ResultadosMultiplicacionMatrices, resultadosNaivKakan...)
	ResultadosMultiplicacionMatrices = append(ResultadosMultiplicacionMatrices, resultadosNaivLoopUnrollingThree...)
	ResultadosMultiplicacionMatrices = append(ResultadosMultiplicacionMatrices, resultadosNaivLoopUnrollingTwo...)
	ResultadosMultiplicacionMatrices = append(ResultadosMultiplicacionMatrices, resultadosNaivLoopUnrollingFour...)
	ResultadosMultiplicacionMatrices = append(ResultadosMultiplicacionMatrices, resultadosWinogradOriginal...)
	ResultadosMultiplicacionMatrices = append(ResultadosMultiplicacionMatrices, resultadosWinogradScaled...)
	ResultadosMultiplicacionMatrices = append(ResultadosMultiplicacionMatrices, resultadosStrassenNaiv...)
	ResultadosMultiplicacionMatrices = append(ResultadosMultiplicacionMatrices, resultadosStrassenWinograd...)
	ResultadosMultiplicacionMatrices = append(ResultadosMultiplicacionMatrices, resultadosIII_sequential_block...)
	ResultadosMultiplicacionMatrices = append(ResultadosMultiplicacionMatrices, resultadosIII_parallel_block...)
	ResultadosMultiplicacionMatrices = append(ResultadosMultiplicacionMatrices, resultadosIV_3_sequential_block...)
	ResultadosMultiplicacionMatrices = append(ResultadosMultiplicacionMatrices, resultadosIV_4_parallel_block...)
	ResultadosMultiplicacionMatrices = append(ResultadosMultiplicacionMatrices, resultadosV_3_Sequential_block...)
	ResultadosMultiplicacionMatrices = append(ResultadosMultiplicacionMatrices, resultadosV_4_parallel_block...)

}

func EscribirResultadoMultiplicacionMatricesTXT() {
	cadena := ""

	for i := range ResultadosMultiplicacionMatrices {
		cadena += ResultadosMultiplicacionMatrices[i].Titulo + " " + strconv.FormatInt(ResultadosMultiplicacionMatrices[i].Duracion, 10) + "\n"
	}

	utilidades.EscribirArchivo("resultados.txt", []byte(cadena))
}

func EscribirResultadoMultiplicacionMatricesJSON() {
	cadena, err := json.Marshal(ResultadosMultiplicacionMatrices)

	utilidades.VerificarError(err)

	utilidades.EscribirArchivo("resultados.json", cadena)
}
