package bmark_multiplicacion_grandes

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/modelos"
	"generador/pkg/tiempo"
)

func BmarkAmericanaRecursivoEstatico(numerosA, numerosB []modelos.NumeroGrande) {
	for i := range numerosA {
		americanaRecursivoEstatico(numerosA[i], numerosB[i])
	}
}

func americanaRecursivoEstatico(a, b modelos.NumeroGrande) {

	r := make([]int, (len(a.Datos)+len(b.Datos))*2)

	defer tiempo.CronometrarMultiplicacionGrandes(modelos.AMERICANA_RECURSIVA_ESTATICO, len(a.Datos))()

	multiplicacion_numeros_grandes.MultiplicacionAmericanaRecursiva(a.Datos, b.Datos, r, len(b.Datos)-1, len(a.Datos)-1, len(r))
}

func BmarkAmericanaRecursivoDinamico(numerosA, numerosB []modelos.NumeroGrande) {
	for i := range numerosA {
		americanaRecursivoDinamico(numerosA[i], numerosB[i])
	}
}

func americanaRecursivoDinamico(a, b modelos.NumeroGrande) {
	l1 := modelos.ListaDesdeSlice(a.Datos)
	l2 := modelos.ListaDesdeSlice(b.Datos)

	r := modelos.Lista((l1.Cantidad + l2.Cantidad) * 2)

	defer tiempo.CronometrarMultiplicacionGrandes(modelos.AMERICANA_RECURSIVA_DINAMICO, len(a.Datos))()

	multiplicacion_numeros_grandes.MultiplicacionAmericanaRecursivaEstructuras(l1, l2, r, l2.Cantidad-1, l1.Cantidad-1, r.Cantidad-1)
}
