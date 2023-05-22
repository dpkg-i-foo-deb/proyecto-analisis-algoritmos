package bmark_multiplicacion_grandes

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkInglesaRecursivoEstatico(numerosA, numerosB []modelos.NumeroGrande) {
	for i := range numerosA {
		inglesaRecursivoEstatico(numerosA[i], numerosB[i])
	}
}

func inglesaRecursivoEstatico(a, b modelos.NumeroGrande) {

	r := make([]int, (len(a.Datos)+len(b.Datos))*2)

	defer tiempo.CronometrarMultiplicacionGrandes(modelos.INGLESA_RECURSIVA_ESTATICO, len(a.Datos))()

	multiplicacion_numeros_grandes.MultiplicacionInglesaRecursiva(a.Datos, b.Datos, r, 0, 0)
}

func BmarkInglesaRecursivoDinamico(numerosA, numerosB []modelos.NumeroGrande) {
	for i := range numerosA {
		inglesaRecursivoDinamico(numerosA[i], numerosB[i])
	}
}

func inglesaRecursivoDinamico(a, b modelos.NumeroGrande) {
	l1 := modelos.ListaDesdeSlice(a.Datos)
	l2 := modelos.ListaDesdeSlice(b.Datos)

	r := modelos.Lista((l1.Cantidad + l2.Cantidad) * 2)

	defer tiempo.CronometrarMultiplicacionGrandes(modelos.INGLESA_RECURSIVA_DINAMICO, len(a.Datos))()

	multiplicacion_numeros_grandes.MultiplicacionInglesaRecursivaEstructuras(l1, l2, r, 0, 0)
}
