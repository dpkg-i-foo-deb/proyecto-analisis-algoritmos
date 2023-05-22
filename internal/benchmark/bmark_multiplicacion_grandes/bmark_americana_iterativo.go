package bmark_multiplicacion_grandes

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkAmericanaIterativaEstatica(numerosA []modelos.NumeroGrande, numerosB []modelos.NumeroGrande) {
	for i := range numerosA {
		americanaIterativaEstatica(numerosA[i], numerosB[i])
	}
}

func americanaIterativaEstatica(a, b modelos.NumeroGrande) {

	r := make([]int, (len(a.Datos)+len(b.Datos))*2)

	defer tiempo.CronometrarMultiplicacionGrandes(modelos.AMERICANA_ITERATIVA_ESTATICO, len(a.Datos))()

	multiplicacion_numeros_grandes.MultiplicacionAmericanaIterativa(a.Datos, b.Datos, r)

}

func BmarkAmericanaIterativaDinamica(numerosA []modelos.NumeroGrande, numerosB []modelos.NumeroGrande) {
	for i := range numerosA {
		americanaIterativaDinamica(numerosA[i], numerosB[i])
	}
}

func americanaIterativaDinamica(a, b modelos.NumeroGrande) {
	l1 := modelos.ListaDesdeSlice(a.Datos)
	l2 := modelos.ListaDesdeSlice(b.Datos)

	r := modelos.Lista((l1.Cantidad + l2.Cantidad) * 2)

	defer tiempo.CronometrarMultiplicacionGrandes(modelos.AMERICANA_ITERATIVA_DINAMICO, len(a.Datos))()

	multiplicacion_numeros_grandes.MultiplicacionAmericanaIterativaEstructuras(l1, l2, r)
}
