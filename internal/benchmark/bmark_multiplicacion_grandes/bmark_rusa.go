package bmark_multiplicacion_grandes

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkRusaIterativa(numerosA, numerosB []modelos.NumeroGrande) {
	for i := range numerosA {
		multiplicarRusaIterativa(numerosA[i], numerosB[i])
	}
}

func multiplicarRusaIterativa(a, b modelos.NumeroGrande) {

	r := make([]int, (len(a.Datos)+len(b.Datos))*2)

	defer tiempo.CronometrarMultiplicacionGrandes(modelos.RUSA_ITERATIVA, len(a.Datos))()

	multiplicacion_numeros_grandes.MultiplicacionRusaIterativa(a.Datos, b.Datos, r)
}

func BmarkRusaRecursiva(numerosA, numerosB []modelos.NumeroGrande) {
	for i := range numerosA {
		multiplicarRusaRecursiva(numerosA[i], numerosB[i])
	}
}

func multiplicarRusaRecursiva(a, b modelos.NumeroGrande) {

	r := make([]int, (len(a.Datos)+len(b.Datos))*2)

	defer tiempo.CronometrarMultiplicacionGrandes(modelos.RUSA_RECURSIVA, len(a.Datos))()

	multiplicacion_numeros_grandes.MultiplicacionRusaRecursiva(a.Datos, b.Datos, r)
}
