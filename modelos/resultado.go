package modelos

import "time"

type AlgoritmoMultuplicacion string

const (
	III_SEQUENTIAL_BLOCK      AlgoritmoMultuplicacion = "III Sequential Block"
	IV_3_SEQUENTIAL_BLOCK     AlgoritmoMultuplicacion = "III 3 Sequential Block"
	NAIV_KAHAN                AlgoritmoMultuplicacion = "NaivKahan"
	NAIV_LOOP_UNROLLING_FOUR  AlgoritmoMultuplicacion = "Naiv Loop Unrolling Four"
	NAIV_LOOP_UNROLLING_THREE AlgoritmoMultuplicacion = "Naiv Loop Unrolling Three"
	NAIV_LOOP_UNROLLING_TWO   AlgoritmoMultuplicacion = "Naiv Loop Unrolling Two"
	NAIV_ON_ARRAY             AlgoritmoMultuplicacion = "Naiv on array"
	NAIV_STANDARD             AlgoritmoMultuplicacion = "Naiv Standard"
	STRASSEN_NAIV             AlgoritmoMultuplicacion = "Strassen Naiv"
	V_3_SEQUENTIAL_BLOCK      AlgoritmoMultuplicacion = "V 3 Sequential Block"
	V_4_PARALLEL_BLOCK        AlgoritmoMultuplicacion = "V 4 Pararell Block"
	WINOGRAD_ORIGINAL         AlgoritmoMultuplicacion = "Winograd Original"
	WINOGRAD_SCALED           AlgoritmoMultuplicacion = "Winograd Scaled"
)

type Resultado struct {
	Titulo    string                  `json:"titulo"`
	Algoritmo AlgoritmoMultuplicacion `json:"algoritmo"`
	N         int                     `json:"n"`
	Duracion  time.Duration           `json:"duracion"`
}
