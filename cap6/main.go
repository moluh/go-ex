package main

import "fmt"

func main() {
	// test1()
	// test2()
	// restricciones()
	// appends()
	// lenCap()
	// Make()
	Copy()
}

func Copy() {
	original := []int{1, 2, 3, 4, 5}
	copia := make([]int, len(original))
	fmt.Println("Original", original)
	fmt.Println("Copia", copia)

	n := copy(copia, original)
	fmt.Println(n, " numeros copiados. Copia: ", copia)

}

func Make() {
	altaCapacidad := make([]float32, 0, 2048)
	fmt.Print("Alta Capacidad", altaCapacidad)
}

func lenCap() {
	var sl []int
	fmt.Printf("longitud %v. capacidad %v\n", len(sl), cap(sl))
	sl = append(sl, 1, 2, 3, 4)
	fmt.Printf("longitud %v. capacidad %v\n", len(sl), cap(sl))
	sl = append(sl, 5)
	fmt.Printf("longitud %v. capacidad %v\n", len(sl), cap(sl))
}

func restricciones() {
	// - El tamañi de un vector es PARTE DEL TIPO. Eso significa que las sig.
	// variables tienen distintos tipos:
	v1 := [3]int{5, 4, 3}
	var v2 [10]int
	fmt.Println("v1", v1)
	fmt.Println("v2", v2)
}

func appends() {
	var p []int
	p = append(p, 1, 2, 3)
	p = append(p, 6)
	fmt.Println("Porción p:", p)
}

func test2() {
	var a []int            // Porción
	b := []int{1, 2, 3}    // Porción
	c := [3]int{1, 2, 3}   // Vector
	d := [...]int{1, 2, 3} // Vector
	fmt.Println("Porción a:", a)
	fmt.Println("Porción b:", b)
	fmt.Println("Vector c:", c)
	fmt.Println("Vector d:", d)
}

func test1() {

	// arr := [5]int{1, 2, 3, 4, 5}
	arr := [...]int{1, 2, 3, 4, 5, 6} // Infiere el compilador, segun la cantidad de datos de asignacion
	// arr := [4]int{1, 2, 3, 4, 5} // <- Eror
	// arr := [7]int{1, 2, 3, 4, 5}
	fmt.Println("arr", arr)

	var tableroAjedrez [8][8]uint8
	fmt.Println("tableroAjedrez", tableroAjedrez)
}
