package utilidades

import (
	"fmt"
	"generador/modelos"
	"generador/resultados"
	"strconv"
	"time"
)

func MedirTiempo(algoritmo modelos.AlgoritmoMultuplicacion, n int) func() time.Duration {
	inicio := time.Now()
	return func() time.Duration {
		resultado := modelos.Resultado{
			Titulo:    string(algoritmo) + " " + strconv.FormatInt(int64(n), 10) + " Elementos",
			Algoritmo: algoritmo,
			N:         n,
			Duracion:  time.Since(inicio).Nanoseconds(),
		}

		resultados.Resultados = append(resultados.Resultados, resultado)

		fmt.Printf("%s %d x %d Tiempo %v\n", algoritmo, n, n, time.Since(inicio))

		return time.Since(inicio)
	}
}
