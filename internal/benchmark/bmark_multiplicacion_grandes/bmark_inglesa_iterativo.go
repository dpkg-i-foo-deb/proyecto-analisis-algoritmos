package bmark_multiplicacion_grandes

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkInglesaIterativaEstatica(numerosA, numerosB []modelos.NumeroGrande) {
	for i := range numerosA {
		inglesaIterativaEstatica(numerosA[i], numerosB[i])
	}
}

func inglesaIterativaEstatica(a, b modelos.NumeroGrande) {
	r := make([]int, (len(a.Datos)+len(b.Datos))*2)

	defer tiempo.CronometrarMultiplicacionGrandes(modelos.INGLESA_ITERATIVA_ESTATICO, len(a.Datos))()

	multiplicacion_numeros_grandes.MultiplicacionInglesaIterativa(a.Datos, b.Datos, r)

}

func BmarkInglesaIterativaDinamica(numerosA, numerosB []modelos.NumeroGrande) {
	for i := range numerosA {
		inglesaIterativaDinamica(numerosA[i], numerosB[i])
	}
}

func inglesaIterativaDinamica(a, b modelos.NumeroGrande) {
	l1 := modelos.ListaDesdeSlice(a.Datos)
	l2 := modelos.ListaDesdeSlice(b.Datos)

	r := modelos.Lista((l1.Cantidad + l2.Cantidad) * 2)

	defer tiempo.CronometrarMultiplicacionGrandes(modelos.INGLESA_ITERATIVA_DINAMICO, len(a.Datos))()

	multiplicacion_numeros_grandes.MultiplicacionInglesaIterativaEstructuras(l1, l2, r)
}
