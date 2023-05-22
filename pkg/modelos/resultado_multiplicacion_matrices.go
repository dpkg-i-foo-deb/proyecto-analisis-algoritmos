package modelos

import "time"

type AlgoritmoMultiplicacionMatrices string

const (
	III_SEQUENTIAL_BLOCK      AlgoritmoMultiplicacionMatrices = "III Sequential Block - Bloque Caché"
	III_PARALLEL_BLOCK        AlgoritmoMultiplicacionMatrices = "III Parallel Block - Bloque Caché"
	IV_3_SEQUENTIAL_BLOCK     AlgoritmoMultiplicacionMatrices = "IV 3 Sequential Block - Bloque = 2"
	IV_4_PARALLEL_BLOCK       AlgoritmoMultiplicacionMatrices = "IV 4 Parallel Block - Bloque = 2"
	NAIV_KAHAN                AlgoritmoMultiplicacionMatrices = "NaivKahan"
	NAIV_LOOP_UNROLLING_FOUR  AlgoritmoMultiplicacionMatrices = "Naiv Loop Unrolling Four"
	NAIV_LOOP_UNROLLING_THREE AlgoritmoMultiplicacionMatrices = "Naiv Loop Unrolling Three"
	NAIV_LOOP_UNROLLING_TWO   AlgoritmoMultiplicacionMatrices = "Naiv Loop Unrolling Two"
	NAIV_ON_ARRAY             AlgoritmoMultiplicacionMatrices = "Naiv on array"
	NAIV_STANDARD             AlgoritmoMultiplicacionMatrices = "Naiv Standard"
	STRASSEN_NAIV             AlgoritmoMultiplicacionMatrices = "Strassen Naiv"
	STRASSEN_WINOGRAD         AlgoritmoMultiplicacionMatrices = "Strassen Winograd"
	V_3_SEQUENTIAL_BLOCK      AlgoritmoMultiplicacionMatrices = "V 3 Sequential Block - Bloque = Matriz"
	V_4_PARALLEL_BLOCK        AlgoritmoMultiplicacionMatrices = "V 4 Pararell Block - Bloque = Doble Matriz"
	WINOGRAD_ORIGINAL         AlgoritmoMultiplicacionMatrices = "Winograd Original"
	WINOGRAD_SCALED           AlgoritmoMultiplicacionMatrices = "Winograd Scaled"
)

type ResultadoAlgoritmoMultiplicacion struct {
	Titulo         string                          `json:"titulo"`
	Algoritmo      AlgoritmoMultiplicacionMatrices `json:"algoritmo"`
	N              int                             `json:"n"`
	Duracion       int64                           `json:"duracion"`
	DuracionHumano time.Duration                   `json:"duracion-humano"`
}
