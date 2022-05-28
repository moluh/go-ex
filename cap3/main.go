package main

import "fmt"

func main() {
	// test1()
	// test2()
	// leyendoModificando()
	leyendoModificando2()
}

func leyendoModificando2() {
	i := 20
	j := 20
	p1 := &i
	p2 := &j

	fmt.Println("============")
	if p1 == p2 {
		fmt.Println("p1 y p2 apuntan a la misma direccion de memoria")
	} else {
		fmt.Println("p1 y p2 apuntan a diferente direccion de memoria")
	}

	if *p1 == *p2 {
		fmt.Println("p1 y p2 apuntan a valores iguales")
	} else {
		fmt.Println("p1 y p2 apuntan a valores diferentes")
	}

	fmt.Println("============")
	fmt.Println("p1", p1)
	fmt.Println("p2", p2)
	fmt.Println("&i", &i)
	fmt.Println("&j", &j)
	fmt.Println("i", i)
	fmt.Println("j", j)
	fmt.Println("*p1", *p1)
	fmt.Println("*p2", *p2)
}

func leyendoModificando() {

	/**
	 &: El operador ampersand (&) retorna la direccion de  memoria de una variable.
	 *: El apuntador se define como una variable, añadiendo un asterisco * delante del
	 	tipo de datos al que este apuntador puede apuntar
		 var pi *int // apuntador a int
	*/

	i := 10
	p := &i
	a := *p
	*p = 21
	fmt.Println("i:", i)
	fmt.Println("p:", p)
	fmt.Println("a:", a)
	fmt.Println("=========:")
	fmt.Println("&i:", &i)
	fmt.Println("*p:", *p)
	fmt.Println("&p:", &p)
	fmt.Println("&a:", &a)
}

func test2() {
	// forma larga
	// i := 10
	// var p *int
	// p = &i
	// fmt.Println("P:", p)

	// forma corta
	i := 10
	p := &i
	fmt.Println("P:", p)
}

func test1() {
	var pi *int
	// pi = nil
	if pi == nil {
		fmt.Println("PI ES NIL:", pi)
	}
}
