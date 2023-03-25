package test

import (
	"generador/utilidades"
	"testing"
)

func TestGeneracionLecturaMatrices(t *testing.T) {
	utilidades.GenerarMatrices()

	matrices := utilidades.LeerMatrices()

	if cap(matrices) < 12 {
		t.Error("Las matrices no son suficientes, fueron generadas ", cap(matrices))
	}
}
