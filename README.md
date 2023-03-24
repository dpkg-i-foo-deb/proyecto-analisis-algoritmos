# proyecto-analisis-algoritmos

## ¿Cómo ejecutar las pruebas?
Ejecutar en la RAIZ del proyecto
`go test ./...`

## Métodos
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


### NaivLoopUnrollingFour

La función `NaivLoopUnrollingFour` toma dos matrices `a` y `b` como entrada y devuelve una matriz resultante que es el resultado de la multiplicación de matrices `a` y `b`. La función utiliza la técnica de desenrollado de bucles para mejorar el rendimiento de la multiplicación de matrices.

1.  Se inicializan las variables `n`, `p`, y `m` con la longitud de la matriz `a`, la cantidad de filas de la matriz `b`, y la longitud de la primera fila de la matriz `b`, respectivamente. También se inicializa una matriz vacía `result` que almacenará el resultado final de la multiplicación.
    
2.  Si la cantidad de filas de la matriz `b` es divisible por 4 y el residuo es 1, se ejecuta el siguiente bloque de código. En este bloque, se itera sobre cada fila `i` de la matriz `a` y se calcula la suma de productos de los elementos correspondientes de la fila `i` de la matriz `a` y las columnas de la matriz `b`. Se utiliza un bucle `for` con incrementos de 4 para procesar 4 columnas de la matriz `b` al mismo tiempo. El resultado se almacena en la variable `suma` y se agrega a una fila auxiliar `filaAux`. Al final de la iteración de cada fila `i`, se agrega la fila `filaAux` a la matriz `result`.
    
3.  Si la cantidad de filas de la matriz `b` es divisible por 4 y el residuo es 2, se ejecuta el siguiente bloque de código. En este bloque, se realiza el mismo proceso que en el bloque anterior, pero se procesan 2 columnas adicionales de la matriz `b` que quedaron por fuera del primer bloque. El resultado se calcula de manera similar a como se hizo en el primer bloque y se agrega a la fila auxiliar `filaAux`.
    
4.  Si la cantidad de filas de la matriz `b` es divisible por 4 y el residuo es 3, se ejecuta el siguiente bloque de código. En este bloque, se procesan 3 columnas adicionales de la matriz `b` que quedaron por fuera de los dos bloques anteriores. El resultado se calcula de manera similar a como se hizo en los bloques anteriores y se agrega a la fila auxiliar `filaAux`.
    
5.  Si la cantidad de filas de la matriz `b` no es divisible por 4, se ejecuta el siguiente bloque de código. En este bloque, se procesan las columnas restantes de la matriz `b` que no fueron procesadas en los bloques anteriores. El resultado se calcula de manera similar a como se hizo en los bloques anteriores y se agrega a la fila auxiliar `filaAux`.
    
6.  Al finalizar el procesamiento de todas las filas de la matriz `a`, la función devuelve la matriz `result` como resultado final de la multiplicación de matrices.
    
En resumen, la función utiliza la técnica de desenrollado de bucles para mejorar el rendimiento de la multiplicación de matrices y se encarga de procesar todas las columnas de la matriz `b` de manera eficiente, dependiendo de la cantidad de columnas que tenga.

### III. Sequential Block

1. Se definen las dos matrices que se desean multiplicar, A y B, y se crea una matriz C de tamaño adecuado para almacenar el resultado.

2. Se define el tamaño del bloque block_size a utilizar en la multiplicación. Este valor se calcula en función del tamaño de la caché del procesador, de manera que cada bloque quepa completamente en la caché para minimizar los accesos a memoria. Por ejemplo, si la caché tiene un tamaño de 8 KB y las matrices son de tamaño 1000x1000, entonces se podría usar un block_size de 125 (ya que 125x125xsizeof(elemento) = 8KB).

3. Se realiza un bucle anidado sobre los bloques de la matriz resultante C. Para cada par de bloques de C que se van a calcular, se realiza un tercer bucle anidado sobre los bloques de las matrices A y B que intervienen en la multiplicación.

4. Para cada par de bloques de A y B que se van a multiplicar, se realiza un bucle anidado sobre las filas y columnas de los bloques. Dentro de este bucle se realiza la multiplicación de la submatriz de A por la submatriz de B y se acumula el resultado en el bloque correspondiente de C.

5. Se repite este proceso para todos los pares de bloques de C, y al final se obtiene la matriz resultante completa.