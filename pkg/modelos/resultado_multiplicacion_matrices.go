package modelos

import "time"

type AlgoritmoMultiplicacion string

const (
	III_SEQUENTIAL_BLOCK      AlgoritmoMultiplicacion = "III Sequential Block - Bloque Caché"
	III_PARALLEL_BLOCK        AlgoritmoMultiplicacion = "III Parallel Block - Bloque Caché"
	IV_3_SEQUENTIAL_BLOCK     AlgoritmoMultiplicacion = "IV 3 Sequential Block - Bloque = 2"
	IV_4_PARALLEL_BLOCK       AlgoritmoMultiplicacion = "IV 4 Parallel Block - Bloque = 2"
	NAIV_KAHAN                AlgoritmoMultiplicacion = "NaivKahan"
	NAIV_LOOP_UNROLLING_FOUR  AlgoritmoMultiplicacion = "Naiv Loop Unrolling Four"
	NAIV_LOOP_UNROLLING_THREE AlgoritmoMultiplicacion = "Naiv Loop Unrolling Three"
	NAIV_LOOP_UNROLLING_TWO   AlgoritmoMultiplicacion = "Naiv Loop Unrolling Two"
	NAIV_ON_ARRAY             AlgoritmoMultiplicacion = "Naiv on array"
	NAIV_STANDARD             AlgoritmoMultiplicacion = "Naiv Standard"
	STRASSEN_NAIV             AlgoritmoMultiplicacion = "Strassen Naiv"
	STRASSEN_WINOGRAD         AlgoritmoMultiplicacion = "Strassen Winograd"
	V_3_SEQUENTIAL_BLOCK      AlgoritmoMultiplicacion = "V 3 Sequential Block - Bloque = Matriz"
	V_4_PARALLEL_BLOCK        AlgoritmoMultiplicacion = "V 4 Pararell Block - Bloque = Doble Matriz"
	WINOGRAD_ORIGINAL         AlgoritmoMultiplicacion = "Winograd Original"
	WINOGRAD_SCALED           AlgoritmoMultiplicacion = "Winograd Scaled"
)

type ResultadoAlgoritmoMultiplicacion struct {
	Titulo         string                  `json:"titulo"`
	Algoritmo      AlgoritmoMultiplicacion `json:"algoritmo"`
	N              int                     `json:"n"`
	Duracion       int64                   `json:"duracion"`
	DuracionHumano time.Duration           `json:"duracion-humano"`
}
