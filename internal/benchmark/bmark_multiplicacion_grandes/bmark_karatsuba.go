package bmark_multiplicacion_grandes

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkKaratsuba(numerosA, numerosB []modelos.NumeroGrande) {
	for i := range numerosA {
		multiplicarKaratsuba(numerosA[i], numerosB[i])
	}
}

func multiplicarKaratsuba(a, b modelos.NumeroGrande) {

	defer tiempo.CronometrarMultiplicacionGrandes(modelos.KARATSUBA, len(a.Datos))()

	multiplicacion_numeros_grandes.MultiplicacionKaratsubaRecursiva(a.Datos, b.Datos)
}
