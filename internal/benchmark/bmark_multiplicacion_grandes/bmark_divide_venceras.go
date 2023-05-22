package bmark_multiplicacion_grandes

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkDivideVenceras(numerosA, numerosB []modelos.NumeroGrande) {
	for i := range numerosA {
		multiplicarDivideVenceras(numerosA[i], numerosB[i])
	}
}

func multiplicarDivideVenceras(a, b modelos.NumeroGrande) {

	r := make([]int, (len(a.Datos)+len(b.Datos))*2)

	defer tiempo.CronometrarMultiplicacionGrandes(modelos.DIVIDE_Y_VENCERAS, len(a.Datos))()

	multiplicacion_numeros_grandes.MultiplicacionDivideVenceras1(a.Datos, b.Datos, r)

}
