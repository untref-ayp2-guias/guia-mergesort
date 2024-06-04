# Guía Ordenamientos avanzados

### Ejercicio 1

Dado un arreglo de números enteros, diseña un algoritmo que encuentre el "k-ésimo" elemento más pequeño del arreglo. Es decir, encuentra el elemento que ocuparía la posición k si el arreglo estuviera ordenado de manera ascendente utilizando el algoritmo QuickSort.

Pasos para resolver el ejercicio:

1. Define una función llamada "EncontrarKesimo" que tome como parámetros el arreglo de números enteros y el valor k.
2. Implementa una lógica similar al particionamiento del algoritmo QuickSort para encontrar el elemento "k-ésimo" en el arreglo.
    1. Seleccione un pivote aleatorio del arreglo.
    2. Realice el particionamiento del arreglo en base al pivote, de manera que todos los elementos menores que el pivote estén a su izquierda y los mayores estén a su derecha.
    3. Verifique en qué posición quedó el pivote después del particionamiento.
    4. Si la posición del pivote es igual a k, significa que hemos encontrado el elemento "k-ésimo" y lo retornamos.
    5. Si la posición del pivote es mayor a k, aplicamos el mismo proceso recursivamente en la sublista izquierda del pivote.
    6. Si la posición del pivote es menor a k, aplicamos el mismo proceso recursivamente en la sublista derecha del pivote.

### Ejercicio 2

### Ejercicio 3

### Ejercicio 4


### Ejercicio 5

1. Implementar el merge sort, pero en vez de dividir en dos partes el arreglo, dividirlo en tres, llamando recursivamente al algoritmo para las tres partes, y luego hacer el merge de las 3.
Suponiendo que el merge de las tres partes es de orden lineal, calcular el orden de este algoritmo utilizando el teorema maestro.

2. Si en vez de separar en 3 partes, se divide en n partes (siendo n la cantidad de elementos del arreglo), ¿a qué otro algoritmo de ordenamiento se asemeja esta implementación? ¿cuál es el orden de dicho algoritmo?

3. Dado lo obtenido en los puntos anteriores, ¿tiene sentido implementar mergesort con separación en k arreglos (para k > 2)?

### Ejercicio 6
### Ejercicio 7
### Ejercicio 8

### Ejercicio 9

Compare las siguientes versiones de HeapSort bajo la luz de la complejidad temporal y espacial:
    1. Aquella que interpreta un array como Heap, y verifica la propiedad de montículo
    2. Aquella que toma cada elemento de un array y lo inserta en un montículo, para luego pasarlos nuevamente al array

### Ejercicio 10

2. Dada una lista de números enteros, se pide encontrar el tercer valor más grande de esa lista utilizando el algoritmo de HeapSort.
### Ejercicio 11
### Ejercicio 12
### Ejercicio 13
### Ejercicio 14
