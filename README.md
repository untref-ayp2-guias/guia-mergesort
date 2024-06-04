# Guía Ordenamientos avanzados

### Ejercicio 1

Implementar quicksort pero con el pivote elegido aleatoriamente. Comparar eficiencia con el dado en clase para un ejemplo con 500 elementos.

### Ejercicio 2

En lápiz y papel dibujar cada paso. Utilizar QuickSort para ordenar el siguiente arreglo de menor a mayor: [10, -5, 3, 0, 1, -42, 13, 10, -8, 9]

### Ejercicio 3
En lápiz y papel dibujar cada paso. Utilizar QuickSort para ordenar el siguiente arreglo de mayor a menor: [11, -3, 7, -2, 15, -92, 88, 13, -8, 9]

### Ejercicio 4

En lápiz y papel dibujar cada paso. Utilizar QuickSort para ordenar el siguiente arreglo de mayor a menor: [13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2,1], eligiendo como pivote el primer elemento del arreglo, ¿Cúal es su complejidad?, realizar el ejercicio nuevamente eligiendo como pivote la media entre 3 elementos del arreglo. ¿Mejoró la complehidad con respecto a elegir el primer elemento como pivote?.

### Ejercicio 5

1. Implementar el MergeSort, pero en vez de dividir en dos partes el arreglo, dividirlo en tres, llamando recursivamente al algoritmo para las tres partes, y luego hacer el merge de las 3.
Suponiendo que el merge de las tres partes es de orden lineal, calcular el orden de este algoritmo utilizando el teorema maestro.

2. Si en vez de separar en 3 partes, se divide en n partes (siendo n la cantidad de elementos del arreglo), ¿a qué otro algoritmo de ordenamiento se asemeja esta implementación? ¿cuál es el orden de dicho algoritmo?

3. Dado lo obtenido en los puntos anteriores, ¿tiene sentido implementar MergeSort con separación en k arreglos (para k > 2)?

### Ejercicio 6
En lápiz y papel dibujar cada paso. Utilizar MergeSort para ordenar el siguiente arreglo de menor a mayor: [10, -5, 3, 0, 1, -42, 13, 10, -8, 9]

### Ejercicio 7
En lápiz y papel dibujar cada paso. Utilizar MergeSort para ordenar el siguiente arreglo de mayor a menor: [11, -3, 7, -2, 15, -92, 88, 13, -8, 9]

### Ejercicio 8

Compare las siguientes versiones de HeapSort bajo la luz de la complejidad temporal y espacial:
    1. Aquella que interpreta un array como Heap, y verifica la propiedad de montículo
    2. Aquella que toma cada elemento de un array y lo inserta en un montículo, para luego pasarlos nuevamente al array

### Ejercicio 9

Dada una lista de números enteros, se pide encontrar el tercer valor más grande de esa lista utilizando el algoritmo de HeapSort.

### Ejercicio 10

En lápiz y papel dibujar cada paso. Utilizar HeapSort para ordenar el siguiente arreglo de menor a mayor: [10, -5, 3, 0, 1, -42, 13, 10, -8, 9]

### Ejercicio 11

En lápiz y papel dibujar cada paso. Utilizar HeapSort para ordenar el siguiente arreglo de mayor a menor: [11, -3, 7, -2, 15, -92, 88, 13, -8, 9]

### Ejercicio 12
Implemenar la función downHeap de RecursiveHeapSort de forma recursiva
