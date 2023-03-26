package resultados

import (
	"encoding/json"
	"generador/modelos"
	"generador/utilidades"
	"os"
	"sort"
	"strconv"
)

var Resultados []modelos.Resultado

var resultadosIII_parallel_block []modelos.Resultado
var resultadosIII_sequential_block []modelos.Resultado
var resultadosIV_3_sequential_block []modelos.Resultado
var resultadosIV_4_parallel_block []modelos.Resultado
var resultadosNaivKakan []modelos.Resultado
var resultadosNaivLoopUnrollingFour []modelos.Resultado
var resultadosNaivLoopUnrollingThree []modelos.Resultado
var resultadosNaivLoopUnrollingTwo []modelos.Resultado
var resultadosNaivOnArray []modelos.Resultado
var resultadosNaivStandard []modelos.Resultado
var resultadosStrassenNaiv []modelos.Resultado
var resultadosStrassenWinograd []modelos.Resultado
var resultadosV_3_Sequential_block []modelos.Resultado
var resultadosV_4_parallel_block []modelos.Resultado
var resultadosWinogradOriginal []modelos.Resultado
var resultadosWinogradScaled []modelos.Resultado

func Consolidar() {
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
	sort.Slice(Resultados, ordenarAscendente(Resultados))

	sort.Slice(resultadosIII_parallel_block, ordenarAscendente(resultadosIII_parallel_block))
	sort.Slice(resultadosIII_sequential_block, ordenarAscendente(resultadosIII_sequential_block))
	sort.Slice(resultadosIV_3_sequential_block, ordenarAscendente(resultadosIV_3_sequential_block))
	sort.Slice(resultadosIV_4_parallel_block, ordenarAscendente(resultadosIV_4_parallel_block))
	sort.Slice(resultadosNaivKakan, ordenarAscendente(resultadosNaivKakan))
	sort.Slice(resultadosNaivLoopUnrollingFour, ordenarAscendente(resultadosNaivLoopUnrollingFour))
	sort.Slice(resultadosNaivLoopUnrollingThree, ordenarAscendente(resultadosNaivLoopUnrollingThree))
	sort.Slice(resultadosNaivLoopUnrollingTwo, ordenarAscendente(resultadosNaivLoopUnrollingTwo))
	sort.Slice(resultadosNaivOnArray, ordenarAscendente(resultadosNaivOnArray))
	sort.Slice(resultadosNaivStandard, ordenarAscendente(resultadosNaivStandard))
	sort.Slice(resultadosStrassenNaiv, ordenarAscendente(resultadosStrassenNaiv))
	sort.Slice(resultadosStrassenWinograd, ordenarAscendente(resultadosStrassenWinograd))
	sort.Slice(resultadosV_3_Sequential_block, ordenarAscendente(resultadosV_3_Sequential_block))
	sort.Slice(resultadosV_4_parallel_block, ordenarAscendente(resultadosV_4_parallel_block))
	sort.Slice(resultadosWinogradOriginal, ordenarAscendente(resultadosWinogradOriginal))
	sort.Slice(resultadosWinogradScaled, ordenarAscendente(resultadosWinogradScaled))

}

func reconstruir() {
	Resultados = []modelos.Resultado{}

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

func ordenarAscendente(arreglo []modelos.Resultado) func(int, int) bool {
	return func(i, j int) bool {
		if arreglo[i].Duracion > arreglo[j].Duracion {
			return arreglo[i].N > arreglo[j].N

		}
		return arreglo[i].N < arreglo[j].N
	}
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

	archivo, err := os.Create("resultados.json")

	utilidades.VerificarError(err)

	defer archivo.Close()

	_, err = archivo.Write(cadena)

	utilidades.VerificarError(err)
}
