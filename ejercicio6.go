/*
Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus servicios.
Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
un arreglo de números enteros con 100 valores
un arreglo de números enteros con 1000 valores
un arreglo de números enteros con 10000 valores

Se debe realizar el ordenamiento de cada una por:
Ordenamiento por inserción
Ordenamiento por burbuja
Ordenamiento por selección

Una go routine por cada ejecución de ordenamiento
Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1000 y después el de 10000.
Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber qué ordenamiento fue mejor para cada arreglo

*/

package main

import (
	"fmt"
	"math/rand"
)

func ordenarBurbuja(miArray []int) []int {
	var aux int

	for i := 0; i < len(miArray); i++ {
		for j := 0; j < len(miArray); j++ {
			if miArray[i] < miArray[j] {
				aux = miArray[i]
				miArray[i] = miArray[j]
				miArray[j] = aux
			}
		}
	}
	return miArray
}

func ordenarInsercion(miArray []int) []int {
	var aux int

	for i := 0; i < len(miArray); i++ {
		aux = miArray[i]
		for j := i - 1; j >= 0 && miArray[j] > aux; j-- {
			miArray[j+1] = miArray[j]
			miArray[j] = aux
		}
	}
	return miArray
}

func ordenarSeleccion(miArray []int) []int {
	var minimo int
	var aux int

	for i := 0; i < len(miArray); i++ {
		minimo = i
		for j := i + 1; j < len(miArray); j++ {
			if miArray[minimo] > miArray[j] {
				minimo = j
			}

			aux = miArray[minimo]
			miArray[minimo] = miArray[i]
			miArray[i] = aux
		}
	}

	return miArray
}

func mostrar(miArray []int) {
	for _, x := range miArray {
		fmt.Printf("%d-", x)
	}
	fmt.Printf("\n")
}

func main() {
	variable1 := rand.Perm(10)
	//variable2 := rand.Perm(1000)
	//variable3 := rand.Perm(10000)

	mostrar(variable1)
	ordenarInsercion(variable1)
	mostrar(variable1)

}
