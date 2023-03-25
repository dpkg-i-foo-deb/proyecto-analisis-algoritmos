package utilidades

import (
	"fmt"
	"strconv"
	"time"
)

func MedirTiempo(algoritmo modelos.AlrogitmoOrdenamiento, cantidadElementos int) func() time.Duration {
	inicio := time.Now()
	return func() time.Duration {
		resultado := modelos.Resultado{
			Titulo:            string(algoritmo) + " " + strconv.FormatInt(int64(cantidadElementos), 10) + " Elementos",
			Algoritmo:         algoritmo,
			CantidadElementos: cantidadElementos,
			Duracion:          time.Since(inicio),
		}

		resultados.Resultados = append(resultados.Resultados, resultado)

		fmt.Printf("%s %d Elementos: Tiempo %v\n", algoritmo, cantidadElementos, time.Since(inicio))

		return time.Since(inicio)
	}
}
