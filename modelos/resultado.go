package modelos

import "time"

type AlgoritmoMultuplication string

const (
	III_SEQUENTIAL_BLOCK      AlgoritmoMultuplication = "III Sequential Block"
	IV_3_SEQUENTIAL_BLOCK     AlgoritmoMultuplication = "III 3 Sequential Block"
	NAIV_KAHAN                AlgoritmoMultuplication = "NaivKahan"
	NAIV_LOOP_UNROLLING_FOUR  AlgoritmoMultuplication = "Naiv Loop Unrolling Four"
	NAIV_LOOP_UNROLLING_THREE AlgoritmoMultuplication = "Naiv Loop Unrolling Three"
	NAIV_LOOP_UNROLLING_TWO   AlgoritmoMultuplication = "Naiv Loop Unrolling Two"
	NAIV_ON_ARRAY             AlgoritmoMultuplication = "Naiv on array"
	NAIV_STANDARD             AlgoritmoMultuplication = "Naiv Standard"
	STRASSEN_NAIV             AlgoritmoMultuplication = "Strassen Naiv"
	V_3_SEQUENTIAL_BLOCK      AlgoritmoMultuplication = "V 3 Sequential Block"
	V_4_PARALLEL_BLOCK        AlgoritmoMultuplication = "V 4 Pararell Block"
	WINOGRAD_ORIGINAL         AlgoritmoMultuplication = "Winograd Original"
	WINOGRAD_SCALED           AlgoritmoMultuplication = "Winograd Scaled"
)

type Resultado struct {
	Titulo    string                  `json:"titulo"`
	Algoritmo AlgoritmoMultuplication `json:"algoritmo"`
	N         int                     `json:"n"`
	Duracion  time.Duration           `json:"duracion"`
}
