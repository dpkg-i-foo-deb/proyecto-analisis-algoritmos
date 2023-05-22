package tiempo

import (
	"fmt"
	"generador/pkg/modelos"
	"generador/pkg/resultados"
	"strconv"
	"time"
)

func CronometrarMultiplicacionMatrices(algoritmo modelos.AlgoritmoMultiplicacionMatrices, n int) func() time.Duration {
	inicio := time.Now()
	return func() time.Duration {
		resultado := modelos.ResultadoMultiplicacionMatrices{
			Titulo:         string(algoritmo) + " " + strconv.FormatInt(int64(n), 10) + " Elementos",
			Algoritmo:      algoritmo,
			N:              n,
			Duracion:       time.Since(inicio).Nanoseconds(),
			DuracionHumano: time.Since(inicio),
		}

		resultados.ResultadosMultiplicacionMatrices = append(resultados.ResultadosMultiplicacionMatrices, resultado)

		fmt.Printf("%s %d x %d Tiempo %v\n", algoritmo, n, n, time.Since(inicio))

		return time.Since(inicio)
	}
}
