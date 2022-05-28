package main

import "fmt"

func main() {

	// max, min, c := MaxMin(20, 30, 32)
	// max, min, c := MaxMinEspecificado(20, 30, 32)
	// fmt.Println("============")
	// fmt.Println("max", max)
	// fmt.Println("min", min)
	// fmt.Println("c", c)

	// identificador vacío, en caso de que no se use el valor retornado
	max, _, _ := MaxMinEspecificado(20, 30, 32)
	fmt.Println("============")
	fmt.Println("max", max)
}

// Funciones retorno de multiples valores
func MaxMin(a, b, c int) (int, int, int) {
	if a > b {
		return a, b, c
	}
	return b, a, c
}

// retorno de valores nombrados
func MaxMinEspecificado(a, b, c int) (max int, min int, C int) {
	if a > b {
		max = a
		min = b
		C = c
	} else {
		max = b
		min = a
		C = c
	}
	return
}
