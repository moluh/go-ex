package main

import "fmt"

func main() {
	// test1()
	// test2()
	signaturaDeFuncion()

	func() {
		fmt.Println("Esto se invocará ")
		fmt.Println("Directamente, pero en otro contexto")
	}()

	/**
	Los literales de funcion tambien pueden recibir el nombre de lambas o "clausuras"(closures)
	*/
}

func test1() {
	var generator func() int

	generator = func() int {
		return 2
	}

	fmt.Println("Generador de 2: ", generator())
}

func test2() {
	// Accediendo fuvera de su ambito:
	count := 0
	var contador func() int
	contador = func() int {
		count++
		return count
	}
	fmt.Println("Contador:  ", contador())
}

func signaturaDeFuncion() {
	var operador func(int, int) int
	operador = Suma
	fmt.Println("Suma = ", operador(3, 4))
	operador = Multiplica
	fmt.Println("Multiplica = ", operador(3, 4))

}

func Multiplica(a, b int) int {
	return a * b
}

func Suma(a, b int) int {
	return a + b
}
