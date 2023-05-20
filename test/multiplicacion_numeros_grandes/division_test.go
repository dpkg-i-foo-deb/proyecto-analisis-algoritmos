package test

import (
	"generador/pkg/algoritmos/multiplicacion_numeros_grandes"
	"generador/pkg/utilidades"
	"testing"
)

func TestDivision(t *testing.T) {
	// Caso de prueba 1: dividendo = 2000, divisor = 2
	dividendo1 := utilidades.FormatearCadenaASlice("2000")
	divisor1 := 2
	esperado1 := "1000"
	resultado1 := utilidades.FormatearSliceACadena(multiplicacion_numeros_grandes.DividirEstatico(dividendo1, divisor1))
	if resultado1 != esperado1 {
		t.Errorf("Error en el caso de prueba 1. Se esperaba '%s' pero se obtuvo '%s'", esperado1, resultado1)
	}

	// Caso de prueba 2: dividendo = 100, divisor = 2
	dividendo2 := utilidades.FormatearCadenaASlice("100")
	divisor2 := 2
	esperado2 := "50"
	resultado2 := utilidades.FormatearSliceACadena(multiplicacion_numeros_grandes.DividirEstatico(dividendo2, divisor2))
	if resultado2 != esperado2 {
		t.Errorf("Error en el caso de prueba 2. Se esperaba '%s' pero se obtuvo '%s'", esperado2, resultado2)
	}

	// Caso de prueba 3: dividendo = 12345678, divisor = 2
	dividendo3 := utilidades.FormatearCadenaASlice("12345678")
	divisor3 := 2
	esperado3 := "6172839"
	resultado3 := utilidades.FormatearSliceACadena(multiplicacion_numeros_grandes.DividirEstatico(dividendo3, divisor3))
	if resultado3 != esperado3 {
		t.Errorf("Error en el caso de prueba 3. Se esperaba '%s' pero se obtuvo '%s'", esperado3, resultado3)
	}

	// Caso de prueba 4: dividendo = 987654321, divisor = 2
	dividendo4 := utilidades.FormatearCadenaASlice("987654321")
	divisor4 := 2
	esperado4 := "493827160"
	resultado4 := utilidades.FormatearSliceACadena(multiplicacion_numeros_grandes.DividirEstatico(dividendo4, divisor4))
	if resultado4 != esperado4 {
		t.Errorf("Error en el caso de prueba 4. Se esperaba '%s' pero se obtuvo '%s'", esperado4, resultado4)
	}

	// Caso de prueba 5: dividendo = 0, divisor = 2
	dividendo5 := utilidades.FormatearCadenaASlice("0")
	divisor5 := 2
	esperado5 := "0"
	resultado5 := utilidades.FormatearSliceACadena(multiplicacion_numeros_grandes.DividirEstatico(dividendo5, divisor5))
	if resultado5 != esperado5 {
		t.Errorf("Error en el caso de prueba 5. Se esperaba '%s' pero se obtuvo '%s'", esperado5, resultado5)
	}
}
