package bmark_multiplicacion_grandes

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkEgipciaIterativo(numerosA, numerosB []modelos.NumeroGrande) {
	for i := range numerosA {
		multiplicarEgipciaIterativo(numerosA[i], numerosB[i])
	}
}

func multiplicarEgipciaIterativo(a, b modelos.NumeroGrande) {

	r := make([]int, (len(a.Datos)+len(b.Datos))*2)

	defer tiempo.CronometrarMultiplicacionGrandes(modelos.EGIPCIA_ITERATIVA, len(a.Datos))()

	multiplicacion_numeros_grandes.MultiplicacionEgipciaIterativa(a.Datos, b.Datos, r)
}

func BmarkEgipciaRecursivo(numerosA, numerosB []modelos.NumeroGrande) {

}

func multiplicarEgipciaRecursivo(a, b modelos.NumeroGrande) {

	r := make([]int, (len(a.Datos)+len(b.Datos))*2)

	defer tiempo.CronometrarMultiplicacionGrandes(modelos.EGIPCIA_RECURSIVA, len(a.Datos))()

	multiplicacion_numeros_grandes.MultiplicacionEgipciaRecursiva(a.Datos, b.Datos, r)
}
