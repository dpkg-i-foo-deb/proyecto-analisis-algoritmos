package bmark_multiplicacion_grandes

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
	"generador/pkg/utilidades"
)

func BmarkCadenas(numerosA, numerosB []modelos.NumeroGrande) {
	for i := range numerosA {
		multiplicarCadenas(numerosA[i], numerosB[i])
	}
}

func multiplicarCadenas(a, b modelos.NumeroGrande) {
	s1 := utilidades.FormatearSliceACadena(a.Datos)
	s2 := utilidades.FormatearSliceACadena(b.Datos)

	defer tiempo.CronometrarMultiplicacionGrandes(modelos.CADENAS, len(a.Datos))()

	multiplicacion_numeros_grandes.MultiplicarUsandoCadenas(s1, s2)

}
