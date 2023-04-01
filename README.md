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

### NaivLoopUnrollingThree
La función `NaivLoopUnrollingThree` desenrosca el ciclo más interno (Aquel que recorre las columnas de la matriz `b`) en tres partes. Esto permite que el código sea más eficiente al reducir el número de iteraciones que se realizan en cada ciclo. La función toma dos matrices `a` y `b` como entrada y devuelve una matriz resultante que es el resultado de la multiplicación de matrices `a` y `b`.

El proceso de la función es el siguiente:

1. Se inicializan los ciclos `for` para recorrer las filas de la matriz `a` y las columnas de la matriz `b`. Se inicializa una matriz vacía `resultado` que almacenará el resultado final de la multiplicación.

2. El tercer ciclo que recorre las columnas de a y las filas de b tendrá un incremento de 3 en 3. Lo que significa que se procesarán 3 columnas de la matriz `b` en cada iteración. Esto se hace para reducir el número de iteraciones que se realizan en cada ciclo.

3. El funcionamiento del ciclo depende si la cantidad de columnas de la matriz `b` es divisible entre 3, en cuyo caso se tendrá en cuenta el último elemento de la matriz `b` en la multiplicación.

4. En cuanto al último elemento se refiere, si `b` es divisible entre 3, se multiplica el último elemento de la fila `i` de la matriz `a` por el último elemento de la columna `j` de la matriz `b` y se suma a la posición `(i,j)` de la matriz `resultado`.

    4.1 Si `b` no es divisible entre 3 y el resultado de la operación módulo es 1, el ciclo que va de k en k se ejecuta hasta la penúltima columna de la matriz `b`. En este caso, se multiplica el último elemento de la fila `i` de la matriz `a` por el último elemento de la columna `j` de la matriz `b` y se suma a la posición `(i,j)` de la matriz `resultado`.

    4.2 Si `b` no es divisible entre 3 y el resultado de la operación módulo es mayor a 1, el ciclo que va de k en k se ejecuta hasta la antepenúltima columna de la matriz `b`. En este caso, se utilizan variables auxiliares para computar los últimos valores de las columnas de `b` con las filas de `a`.

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

### III. Parallel Block

El método `III Parallel Block` utiliza la multiplicación de matrices por bloques de manera similar a `III Sequential Block` pero crea una gorutina por cada bloque de multiplicación que se realiza. Esto permite que la multiplicación de matrices se realice de manera paralela. El tamaño de bloque se calcula teniendo en cuenta el tamaño de la caché L1 del procesador, además debe ser un divisor de las columnas de a y las filas de b. El proceso se puede resumir al siguiente algoritmo:

1. Se definen las dos matrices que se desean multiplicar, A y B, y se crea una matriz C de tamaño adecuado para almacenar el resultado.

2. Se define el tamaño del bloque teniendo en cuenta la caché del procesador y las dimensiones de las matrices a y b.

3. Se realiza una multiplicación por bloques similar a la vista en el método `III Sequential Block` pero en este caso se crea una gorutina por cada bloque de multiplicación que se realiza.

4. Se sincronizan los índices de las matrices a y b para que se realice la multiplicación de matrices de manera correcta.

5. Se espera a que cada gorutina finalice su ejecución.

6. Se retorna el resultado.

Este método es extremadamente eficiente ya que el tamaño de bloque permite al procesador realizar las operaciones directamente en la memoria caché (Esto sin tener en cuenta el sistema operativo u otras aplicaciones en ejecución) 

### Strassen Naiv
Strassen Naiv tiene realiza la multiplicación de matrices utilizando recursividad y la fórmula de Strassen, este método dice que la multiplicación de matrices se puede realizar de la siguiente manera:

    ``C = A * B

    C11 = A11 * B11 + A12 * B21
    C12 = A11 * B12 + A12 * B22
    C21 = A21 * B11 + A22 * B21
    C22 = A21 * B12 + A22 * B22``

Donde A11, A12, A21, A22, B11, B12, B21, B22 son submatrices de A y B. En palabras más simples, los cuadrantes superior izquierdo, superior derecho, inferior izquierdo e inferior derecho de las matrices a y b

El algoritmo se resume a:

1. Si las matrices son de tamaño 1x1, se realiza la multiplicación de matrices de manera normal. Cabe aclarar que siguiendo el algoritmo de Strassen de manera purista, deberíamos seguir este camino. Sin embargo, en la práctica, el algoritmo de Strassen es más lento que el algoritmo de multiplicación de matrices tradicional para matrices pequeñas. Por lo tanto, se utiliza el algoritmo de multiplicación de matrices tradicional `NaivStandard` para matrices pequeñas. En nuestro caso, aquellas de 16x16 o menos.

2. Si las matrices son de tamaño mayor a 16x16, se divide la matriz de entrada A en cuatro submatrices de igual tamaño: A11, A12, A21 y A22. Se divide la matriz de entrada B en cuatro submatrices de igual tamaño: B11, B12, B21 y B22. Cada submatriz tiene la mitad del tamaño de la matriz original.

3. Se encuentran siete productos m1, m2, m3, m4, m5, m6 y m7 utilizando las siguientes fórmulas:

    ``m1 = (A11 + A22) * (B11 + B22)
    m2 = (A21 + A22) * B11
    m3 = A11 * (B12 - B22)
    m4 = A22 * (B21 - B11)
    m5 = (A11 + A12) * B22
    m6 = (A21 - A11) * (B11 + B12)
    m7 = (A12 - A22) * (B21 + B22)``

    3.1 Para encontrar los productos es necesario realizar la multiplicación de sumas y restas de los cuadrantes de las matrices a y b, lo que significa que debemos llamar el algoritmo de Strassen para encontrar los productos m1, m2, m3, m4, m5, m6 y m7. Generando más llamadas recursivas.

4. Una vez encontrados los productos m1, m2, m3, m4, m5, m6 y m7, se calculan los cuadrantes de la matriz resultante C utilizando las siguientes fórmulas:

    ``C11 = m1 + m4 - m5 + m7
    C12 = m3 + m5
    C21 = m2 + m4
    C22 = m1 - m2 + m3 + m6``

5. Se construye una matriz resultado utilizando los cuadrantes C11, C12, C21 y C22 y es retornada.


### Strassen Winograd
1. División de las matrices en submatrices:
Se divide la matriz de entrada A en cuatro submatrices de igual tamaño: A11, A12, A21 y A22. Se divide la matriz de entrada B en cuatro submatrices de igual tamaño: B11, B12, B21 y B22. Cada submatriz tiene la mitad del tamaño de la matriz original.

2. Cálculo de siete productos entre submatrices:
Se calculan siete productos entre submatrices de las matrices A y B, utilizando las siguientes fórmulas:

P1 = A11 * (B12 - B22)
P2 = (A11 + A12) * B22
P3 = (A21 + A22) * B11
P4 = A22 * (B21 - B11)
P5 = (A11 + A22) * (B11 + B22)
P6 = (A12 - A22) * (B21 + B22)
P7 = (A11 - A21) * (B11 + B12)

Estas operaciones permiten calcular el producto de las matrices A y B con solo siete productos de submatrices en lugar de los ocho requeridos por el algoritmo estándar de multiplicación de matrices.

3. Cálculo de las submatrices de la matriz de salida:
Se calculan las cuatro submatrices de la matriz de salida C utilizando las siguientes fórmulas:

C11 = P5 + P4 - P2 + P6
C12 = P1 + P2
C21 = P3 + P4
C22 = P5 + P1 - P3 - P7

Estas fórmulas combinan los siete productos de submatrices calculados en el paso anterior para obtener las submatrices de la matriz de salida.

4. Unión de las submatrices:
Las cuatro submatrices de la matriz de salida C se unen para formar la matriz de salida completa.
    a. C11 = P5 + P4 - P2 + P6:
    En esta operación, se suman los productos P5, P4 y P6 y se le resta el producto P2. El resultado se almacena en la submatriz C11 de la matriz de salida.

    b. C12 = P1 + P2:
    En esta operación, se suman los productos P1 y P2 y se almacena el resultado en la submatriz C12 de la matriz de salida.

    c. C21 = P3 + P4:
    En esta operación, se suman los productos P3 y P4 y se almacena el resultado en la submatriz C21 de la matriz de salida.

    d. C22 = P5 + P1 - P3 - P7:
    En esta operación, se suman los productos P5 y P1, y se les resta los productos P3 y P7. El resultado se almacena en la submatriz C22 de la matriz de salida.

En resumen, el algoritmo de multiplicación de matrices Strassen-Winograd divide las matrices de entrada en submatrices más pequeñas, realiza siete productos de submatrices, combina los resultados para obtener las submatrices de la matriz de salida y luego une las submatrices para formar la matriz de salida completa.

### WinogradOiriginal
El método __`WinogradOriginal`__ se trata de una variante del algoritmo de Straseen que utiliza menos operaciones de multiplicación y suma que el algoritmo clásico.

`WinogradOriginal`__ multiplica dos matrices A y B de tamaño MxP y PxN respectivamente y consta de los siguientes pasos:

1. Se definen las dimensiones de las matrices de entrada A y B
Donde N (columnas de B), P (filas de B) y M (filas de A).

2. Se crea una matriz Resultado de tamaño MxN 
Para almacenar el resultado final de la multiplicación.

3. Se define una variable auxiliar, i, j y k 
Para los índices de los bucles.

4. Se calcula el valor de "upsilon"
Que es la cantidad de filas de B que quedan fuera de los cálculos auxiliares.

5. Se define gamma 
Como la cantidad de filas de B que se usan en los cálculos auxiliares.

6. Se crea un vector "y" de tamaño M y un vector "z" de tamaño N 
Para almacenar los resultados de los cálculos auxiliares.

7. Se realiza un bucle para cada columna i de la matriz B, en el que se calcula el valor de "z[i]" como la suma de los productos de los elementos de las filas pares e impares de la columna i de B.

8. Se realiza un bucle para cada columna i de la matriz B, en el que se calcula el valor de z[i] como la suma de los productos de los elementos de las filas pares e impares de la columna i de B.

9. Si upsilon es igual a 1, entonces P es impar
Se realiza un bucle anidado para calcular el resultado final de la multiplicación. Para cada par de índices i y k, se calcula aux como la suma de los productos de los elementos de las filas pares e impares de la fila i de A y la columna k de B, sumando el elemento A[i][P] de la fila i de A y el elemento B[P][k] de la columna k de B. Luego, se resta y[i] y z[k] y se agrega el resultado a Resultado[i][k].

10. Si upsilon es diferente de 1, entonces P es par
Se realiza un bucle anidado para calcular el resultado final de la multiplicación. Para cada par de índices i y k, se calcula aux como la suma de los productos de los elementos de las filas pares e impares de la fila i de A y la columna k de B, y se resta y[i] y z[k]. Luego, se agrega el resultado a Resultado[i][k].

11.Se devuelve la matriz Resultado como el resultado final de la multiplicación de las matrices A y B.

Este algoritmo es especialmente útil cuando ambas matrices de entrada son de gran tamaño.

### WinogradScaled
El método __`WinogradScaled`__ se basa en la técnica de multiplicación de matrices _'WinogradOriginal'_

La diferencia radica en que antes de llamar dicho algoritmo, se escalan las matrices mediante la función _'MultiplicarEscalar'_ para mejorar la eficiencia de este método.

El algoritmo inicia obteniendo las dimensiones de las matrices de entrada, creando copias escaladas de ambas matrices y calculando los factores de escala lambda, a y b. 

1. ¿Cómo se calcula lambda? 
lambda = floor(0.5 + log(b/a)/log(4))

Donde "a" y "b" son los factores de normalización de las matrices de entrada A y B, respectivamente.

Para obtener estos factores de normalización, se utiliza la función _'NormInf'_, que calcula la norma infinita de la matriz, es decir, el valor absoluto máximo de todos los elementos de la matriz.

Una vez que se tienen los factores de normalización a y b, se calcula el valor de lambda utilizando la fórmula mencionada anteriormente. El valor obtenido de lambda se utiliza para ajustar el tamaño de ambas matrices mediante la función _'MultiplicarEscalar'_.

Posteriormente, se llama a la función WinogradOriginal con las matrices escaladas, que devuelve la matriz resultado. 

Finalmente, se devuelve el resultado de la multiplicación de matrices escaladas.

Finalmente, algoritmo WinogradScaled utiliza la técnica WinogradOriginal para multiplicar matrices grandes y la mejora mediante la escala de matrices de entrada para reducir el costo computacional del algoritmo.

### IV.3 Sequential block
El método __`IV.3 Sequential block`__ implementa la multiplicación de matrices grandes utilizando el enfoque de bloques secuenciales. En este caso, el tamaño de bloque es fijo y se define en el código fuente.

Los pasos presentes en este método son los siguientes: 

1. Se define el tamaño de las matrices de entrada A y B.

2. Se define el tamaño del bloque que se utilizará para dividir las matrices de entrada. En nuestro caso, está definido en 2.

3. Se crea la matriz de salida C, inicializada con ceros y del mismo tamaño que las matrices de entrada.

4. Se inicia el bucle externo que recorre las filas de A y C en bloques de tamaño bsize.

5. Dentro del bucle externo, se inicia un bucle que recorre las columnas de B y C en bloques de tamaño bsize.

    6. En este se inicia un bucle que recorre las columnas de A y las filas de B en bloques de tamaño bsize.

        7. Luego, se inicia un bucle que recorre las filas de A y las columnas de B en los bloques definidos en los bucles anteriores.

8. Finalmente, retorna "C" la cual corresponde a la matriz resultante de la multiplicación.

Este método puede tener un rendimiento variable comparado con `III. Sequential Block` ya que el tamaño de bloque es fijo y podría no aprovechar el rendimiento de la CPU.

### IV.4 Parallel Block
El método `IV.4 Parallel Block` es el mismo que `III. Parallel Block`. Sin embargo, el tamaño de bloque es 2. Esto significa que para una matriz de 4096x4096, se dividirá en 2048 bloques de 2x2. En otras palabras, se crearán 2048 hilos para realizar la multiplicación de matrices. Esto no sólo significa que el procesador estará sobreutilizado, sino que también se desperdiciarán recursos de memoria y caché en operaciones que resultan demasiado pequeñas y repetitivas para aprovechar el rendimiento de la CPU. El algoritmo se resume a:

1. Se define el tamaño de las matrices de entrada A y B.

2. Se define el tamaño del bloque que se utilizará para dividir las matrices de entrada. En nuestro caso, está definido en 2.

3. Se crea la matriz de salida C, inicializada con ceros y del mismo tamaño que las matrices de entrada.

4. Se inician los ciclos externos que recorrerán las matrices a y b en bloques de tamaño 2.

5. Se crea una gorutina por cada bloque de tamaño 2x2.

6. Se sincronizan los índices de los ciclos externos y se espera a que todas las gorutinas finalicen su ejecución.

7. Se retorna la matriz C.

Este método es extremadamente ineficiente y no debe utilizarse para multiplicar matrices grandes. Esto gracias a la sobrecarga de procesador y memoria caché que causa la creación de tantos hilos. En la práctica es incluso más lento que `NaivStandard`

### V.3 Sequential block
El método __`V.3 Sequential block`__ es similar al agoritmo __`IV.3 Sequential block`__ pero con una transposición de las matrices A y B. Es decir, en lugar de multiplicar la fila i de A por la columna j de B, se multiplicará la fila k de A por la columna j de B, y el resultado se almacenará en la posición (k, i) de la matriz resultante.

1.Se define la dimensión de las matrices A y B y el tamaño del bloque (bsize) que se utilizará para dividir la matriz en submatrices más pequeñas.

2.Se crea una matriz C del mismo tamaño que las matrices A y B, y se inicializa con ceros.

3.Se inicia un bucle que se ejecutará para cada bloque de tamaño bsize. El bucle comienza por los índices i1, j1, y k1.

4.Se inicia otro bucle anidado que recorre los elementos de la submatriz i1 a i1+bsize de la matriz A, desde el índice i hasta el índice ObtenerMinimo(i1+bsize, size).

    La función ObtenerMinimo se usa para determinar la cantidad de filas y columnas de la matriz que se deben procesar en cada iteración de los bucles anidados que se utilizan para realizar la multiplicación de matrices en bloques.

    La utilidad de esta función radica en que asegura que el algoritmo de multiplicación de matrices en bloques procese el menor número posible de filas y columnas en cada iteración de los bucles anidados, lo que aumenta la eficiencia del algoritmo y reduce el uso de recursos del sistema.

5.Se inicia otro bucle anidado que recorre los elementos de la submatriz j1 a j1+bsize de la matriz B, desde el índice j hasta el índice ObtenerMinimo(j1+bsize, size).

6.Se inicia otro bucle anidado que recorre los elementos de la submatriz k1 a k1+bsize de la matriz A transpuesta (A traspuesta), desde el índice k hasta el índice ObtenerMinimo(k1+bsize, size).

7.Se realiza la multiplicación de los elementos A[k][j] y B[j][i], y se suma el resultado a la posición C[k][i].

8.El algoritmo termina y se devuelve la matriz C resultante.

### V.4 Parallel block
El método __`V.4 Parallel`__usa como idea principal la división de la tarea en múltiples subprocesos que puedan ejecutarse en paralelo, reduciendo el tiempo de ejecución total del algoritmo. En este caso implementamos un tamaño de bloque igual al doble del tamaño de las matrices

Los pasos presentes en dicho método son los siguientes: 

1. Se obtiene el tamaño de la matriz y se establece un tamaño de bloque predeterminado para la multiplicación de matrices.

2. Se inicializa la matriz C que contendrá el resultado de la multiplicación de matrices.

3. Se crea una variable "wg" del tipo sync.WaitGroup para sincronizar la finalización de todos los subprocesos creados.

4. Se establece la cantidad de subprocesos necesarios para recorrer la matriz de entrada A y B, la cual es igual a "-size / bsize * size / bsize". Esto se realiza para dividir la matriz en bloques de tamaño "bsize x bsize".

5. Se recorre la matriz de entrada A y B por bloques. Por cada bloque (i1,j1) se crea un nuevo subproceso que recorrerá todos los elementos de la matriz C que están dentro del bloque (i1,j1). El subproceso calcula el producto punto a punto de la sección de la matriz A y la sección de la matriz B correspondiente y agrega el resultado al elemento correspondiente en la matriz C.

6. Cada subproceso se agrega a la variable "wg" para esperar su finalización antes de devolver la matriz C.

7. Se espera a que todos los subprocesos finalicen utilizando "wg.Wait()".

8. Se retorna la matriz C con el resultado de la multiplicación de matrices.

Este algoritmo es extremadamente ineficiente, al tener un tamaño de bloque tan grande, no se crearán las suficientes gorutinas para aprovechar el paralelismo de la computadora, por lo que el algoritmo se ejecutará de forma secuencial llegando a tardar incluso más tiempo que `NaivStandard` en ejecutarse.
