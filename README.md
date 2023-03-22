# proyecto-analisis-algoritmos

## ¿Cómo ejecutar las pruebas?
Ejecutar en la RAIZ del proyecto
`go test ./...`

## Métodos

### NaivStandard
El método __`NaivStandard`__ obedece al algoritmo de multiplicación de matrices que hemos utilizado
toda la vida, recibe dos matrices a y b de números enteros y retorna la matriz resultado, es bastante
intuitivo de entender pero igualmente ineficiente.

El algoritmo funciona de la siguiente manera:

1. Se crea una matriz vacía llamada __`resultado`__.

2. Se realiza un recorrido de filas y columnas tradicional para almacenar los resultados en la matriz resultado.

3. Se genera un tercer ciclo durante el recorrido de filas y columnas que se encarga de multiplicar las columnas de a por las filas de b.

4. El resultado de cada multiplicación realizada en el tercer ciclo es almacenado de manera temporal en una variable auxiliar llamada acumulador.

5. Al finalizar el tercer ciclo, se almacena el valor del acumulador en la posición i,j de la matriz resultado.

6. El proceso de repite hasta que el recorrido de filas y columnas ha finalizado.

### NaivOnArray
El método __`NaivOnArray`__ tiene exactamente el mismo recorrido visto en __`NaivStandard`__. Sin embargo, es un poco más eficiente en el uso de recursos al no utilizar una variable extra como acumulador. En su lugar, almacena la productoria de las filas de a y las columnas de b directamente en la matriz resultado. A pesar de esto, la implementación sigue teniendo un orden de complejidad de O(n^3)

### NaivKahan
El método __`NaivKahan`__ nuevamente posee el mismo comportamiento que __`NaivStandard`__ en el recorrido de las matrices. Sin embargo, emplea la sumatoria de Kahan para almacenar el resultado de la productoria. Este método ayuda a reducir la inestabilidad numérica al momento de las sumas.

Es necesario tener en cuenta lo siguiente:

1. Se declaran variables extra para almacenar el resultado de la suma y el error.

2. A través de la suma de Kahan, se cancela el error de la suma en la siguiente iteración, corrigiendo así posibles defectos en las operaciones.

A pesar de incrementar la estabilidad numérica de la productoria, este método sigue teniendo la complejidad de O(n^3) de las implementaciones tradicionales de multiplicación de matrices.

### NaivLoopUnrollingTwo
El método __`NaivLoopUnrollingTwo`__ recibe dos matrices a y b de enteros y retorna la multiplicación de ambas matrices como una nueva matriz de enteros. Este método utiliza una técnica de optimización llamada "loop unrolling" para mejorar el rendimiento de la multiplicación de matrices.

El algoritmo funciona de la siguiente manera:

1. Se crea una matriz vacía llamada __`resultado`__.

2. Si la cantidad de columnas de __`b`__ es par, se realiza el siguiente proceso:

    2.1. Se recorre cada fila __`i`__ de __`a`__.
    
    2.2. Para cada fila __`i`__ de __`a`__, se crea una fila vacía llamada __`filaAux`__ en __`resultado`__.
    
    2.3. Se recorre cada columna __`j`__ de __`b`__.

    2.4. Para cada columna __`j`__ de __`b`__, se realiza la siguiente operación:

        2.4.1. Se crea una variable __`suma`__ inicializada en 0.

        2.4.2. Se recorre cada columna __`k`__ de __`b`__ de dos en dos.

        2.4.3. Para cada columna __`k`__ de __`b`__, se realiza la siguiente operación:

            a. Se multiplica __`a[i][k]`__ por __`b[k][j]`__ y se suma a __`suma`__.

            b. Se multiplica __`a[i][k+1]`__ por __`b[k+1][j]`__ y se suma a __`suma`__.

        2.4.4. El resultado de la suma se guarda en la posición __`(i,j)`__ de __`filaAux`__.

    2.5. La fila __`filaAux`__ se agrega al final de __`resultado`__.
    
3. Si la cantidad de columnas de __`b`__ es impar, se realiza el siguiente proceso:
    
    3.1. Se calcula el índice del último elemento de __`b`__ como __`ultimoIndice`__.
    
    3.2. Se recorre cada fila __`i`__ de __`a`__.
    
    3.3. Para cada fila __`i`__ de __`a`__, se crea una fila vacía llamada __`filaAux`__ en __`resultado`__.

    3.4. Se recorre cada columna __`j`__ de __`b`__.
    
    3.5. Para cada columna __`j`__ de __`b`__, se realiza la siguiente operación:
        
        3.5.1. Se crea una variable __`suma`__ inicializada en 0.
        3.5.2. Se recorre cada columna __`k`__ de __`b`__ de dos en dos hasta ultimoIndice.

        3.5.3. Para cada columna __`k`__ de __`b`__, se realiza la siguiente operación:
            a. Se multiplica __`a[i][k]`__ por __`b[k][j]`__ y se suma a __`suma`__.
            b. Se multiplica __`a[i][k+1]`__ por __`b[k+1][j]`__ y se suma a __`suma`__.
        3.5.4. El resultado de la suma se guarda en la posición __`(i,j)`__ de __`filaAux`__.
    3.6. Se multiplica el último elemento de __`a[i]`__ por el último elemento de __`b`__ y se suma a la posición __`(i,j)`__ de __`filaAux`__.

    3.7. La fila __`filaAux`__ se agrega al final de __`resultado`__.
    
4. Se retorna __`resultado`__.

## III. Sequential Block
1. Se definen las dos matrices que se desean multiplicar, A y B, y se crea una matriz C de tamaño adecuado para almacenar el resultado.

2. Se define el tamaño del bloque block_size a utilizar en la multiplicación. Este valor se calcula en función del tamaño de la caché del procesador, de manera que cada bloque quepa completamente en la caché para minimizar los accesos a memoria. Por ejemplo, si la caché tiene un tamaño de 8 KB y las matrices son de tamaño 1000x1000, entonces se podría usar un block_size de 125 (ya que 125x125xsizeof(elemento) = 8KB).

3. Se realiza un bucle anidado sobre los bloques de la matriz resultante C. Para cada par de bloques de C que se van a calcular, se realiza un tercer bucle anidado sobre los bloques de las matrices A y B que intervienen en la multiplicación.

4. Para cada par de bloques de A y B que se van a multiplicar, se realiza un bucle anidado sobre las filas y columnas de los bloques. Dentro de este bucle se realiza la multiplicación de la submatriz de A por la submatriz de B y se acumula el resultado en el bloque correspondiente de C.

5. Se repite este proceso para todos los pares de bloques de C, y al final se obtiene la matriz resultante completa.