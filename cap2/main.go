package main

import "fmt"

func main() {
	// numbers()
	// text()
	// conversiones()
	constantes()
}

func numbers() {
	/**
	VARIABLES NUMERICAS
	*/
	var minutos int
	minutos = 20 // si no se define, valor por defecto es 0
	var segundos int64 = 2312312533234234234
	var meses int = 12
	anos := 27

	fmt.Println("Minutos:", minutos)
	fmt.Println("Segundos:", segundos)
	fmt.Println("Meses:", meses)
	fmt.Println("Tengo años:", anos)
}

func text() {

	/**
	VARIABLES TEXTO
	*/
	var text string
	text = "Primer texto" // si no se define, valor por defecto es 0
	var text2 string = "\n- Hola, ¿Cómo estás \"Chapatin\"? \n - Yo re bien, o te cuento?"

	text3 := `
	- Hola, ¿Cómo estás "Chapatin"?
	- Yo re bien, o te cuento?
	`

	fmt.Println("text1:\n======\n", text)
	fmt.Println("text2:\n======\n", text2)
	fmt.Println("text3:\n======\n", text3)
}

func conversiones() {
	distancia := 21.23 // float64
	kms := int(distancia)
	kmsFloat := float32(distancia)
	fmt.Println("distancia:\n========================\n", distancia)
	fmt.Println("kms:\n========================\n", kms)
	fmt.Println("kmsFloat:\n========================\n", kmsFloat)
}

func constantes() {
	// const Pi = 3.1416
	const Pi float64 = 3.1416

	const (
		TipoFuente   = "Times New Roman"
		AlturaFuente = 12
		Subrayado    = false
		Negrita      = true
	)
	fmt.Println("Pi:\n========================\n", Pi)
	fmt.Println("TipoFuente:\n========================\n", TipoFuente)
	fmt.Println("AlturaFuente:\n========================\n", AlturaFuente)
	fmt.Println("Subrayado:\n========================\n", Subrayado)
	fmt.Println("Negrita:\n========================\n", Negrita)
}
