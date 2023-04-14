package tiempo

import (
	"fmt"
	"generador/pkg/modelos"
	"generador/resultados"
	"strconv"
	"time"
)

func MedirTiempo(algoritmo modelos.AlgoritmoMultiplicacion, n int) func() time.Duration {
	inicio := time.Now()
	return func() time.Duration {
		resultado := modelos.Resultado{
			Titulo:         string(algoritmo) + " " + strconv.FormatInt(int64(n), 10) + " Elementos",
			Algoritmo:      algoritmo,
			N:              n,
			Duracion:       time.Since(inicio).Nanoseconds(),
			DuracionHumano: time.Since(inicio),
		}

		resultados.Resultados = append(resultados.Resultados, resultado)

		fmt.Printf("%s %d x %d Tiempo %v\n", algoritmo, n, n, time.Since(inicio))

		return time.Since(inicio)
	}
}
