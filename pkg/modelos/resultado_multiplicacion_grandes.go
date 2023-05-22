package modelos

import "time"

type AlgoritmoMultiplicacionNumerosGrandes string

const (
	AMERICANA_ITERATIVA_ESTATICO AlgoritmoMultiplicacionNumerosGrandes = "Multiplicación Americana Iterativa Estática"
	AMERICANA_ITERATIVA_DINAMICO AlgoritmoMultiplicacionNumerosGrandes = "Multiplicación Americana Iterativa Dinámica"
	AMERICANA_RECURSIVA_ESTATICO AlgoritmoMultiplicacionNumerosGrandes = "Multiplicación Americana Recursiva Estática"
	AMERICANA_RECURSIVA_DINAMICO AlgoritmoMultiplicacionNumerosGrandes = "Multiplicación Americana Recursiva Dinámica"
	INGLESA_ITERATIVA_ESTATICO   AlgoritmoMultiplicacionNumerosGrandes = "Multilpicación Inglesa Iterativa Estática"
	INGLESA_ITERATIVA_DINAMICO   AlgoritmoMultiplicacionNumerosGrandes = "Multiplicación Inglesa Iterativa Dinámica"
	INGLESA_RECURSIVA_ESTATICO   AlgoritmoMultiplicacionNumerosGrandes = "Multiplicación Inglesa Recursiva Estática"
	INGLESA_RECURSIVA_DINAMICO   AlgoritmoMultiplicacionNumerosGrandes = "Multiplicación Inglesa Recursiva Dinámica"
	RUSA_ITERATIVA               AlgoritmoMultiplicacionNumerosGrandes = "Multiplicación Rusa Iterativa"
	RUSA_RECURSIVA               AlgoritmoMultiplicacionNumerosGrandes = "Multiplicación Rusa Recursiva"
	HINDU_ITERATIVA              AlgoritmoMultiplicacionNumerosGrandes = "Multiplicación Hindú Iterativa"
	HINDU_RECURSIVA              AlgoritmoMultiplicacionNumerosGrandes = "Multiplicación Hindú Recursiva"
	EGIPCIA_ITERATIVA            AlgoritmoMultiplicacionNumerosGrandes = "Multipicación Egipcia Estática"
	EGIPCIA_RECURSIVA            AlgoritmoMultiplicacionNumerosGrandes = "Multipicación Egipcia Recursiva"
	KARATSUBA                    AlgoritmoMultiplicacionNumerosGrandes = "Karatsuba Recursivo"
	CADENAS                      AlgoritmoMultiplicacionNumerosGrandes = "Multiplicación Representada con Cadenas"
	DIVIDE_Y_VENCERAS            AlgoritmoMultiplicacionNumerosGrandes = "Multiplicación Utilizando Descomposiciones"
)

type ResultadoMultiplicacionNumerosGrandes struct {
	Titulo         string                                `json:"titulo"`
	Algoritmo      AlgoritmoMultiplicacionNumerosGrandes `json:"algoritmo"`
	N              int                                   `json:"n"`
	Duracion       int64                                 `json:"duracion"`
	DuracionHumano time.Duration                         `json:"duracion-humano"`
}
