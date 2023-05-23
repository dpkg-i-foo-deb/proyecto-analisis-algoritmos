package bmark_multiplicacion_grandes

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkHinduIterativa(numerosA, numerosB []modelos.NumeroGrande) {
	for i := range numerosA {
		multiplicarHinduIterativo(numerosA[i], numerosB[i])
	}
}

func multiplicarHinduIterativo(a, b modelos.NumeroGrande) {

	r := make([]int, (len(a.Datos)+len(b.Datos))*2)

	defer tiempo.CronometrarMultiplicacionGrandes(modelos.HINDU_ITERATIVA, len(a.Datos))()

	multiplicacion_numeros_grandes.MultiplicacionHinduIterativa(a.Datos, b.Datos, r)
}

func BmarkHinduRecursiva(numerosA, numerosB []modelos.NumeroGrande) {
	for i := range numerosA {
		multiplicarHinduRecursiva(numerosA[i], numerosB[i])
	}
}

func multiplicarHinduRecursiva(a, b modelos.NumeroGrande) {
	r := make([]int, (len(a.Datos)+len(b.Datos))*2)

	defer tiempo.CronometrarMultiplicacionGrandes(modelos.HINDU_RECURSIVA, len(a.Datos))()

	multiplicacion_numeros_grandes.MultiplicacionHinduRecursiva(a.Datos, b.Datos, r, 0, 0)
}
