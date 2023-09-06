package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.141516

	fmt.Println("pi es:", pi)
	fmt.Println("Pi2 es:", pi2)

	// declaración de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	//Area del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("El area del cuadrado es:", areaCuadrado)

	x := 10
	y := 50

	//suma
	result := x + y
	fmt.Println("El resultado de la suma es:", result)

	//resta
	result = y - x
	fmt.Println("El resultado de la resta es:", result)

	// multiplicación
	result = x * y
	fmt.Println("El resultado de la multiplicación es:", result)

	//división
	result = y / x
	fmt.Println("El resultado de laa división es:", result)

	//modulo ó residuo
	result = y % x
	fmt.Println("El modulo es:", result)

	//Incrementar
	x++
	fmt.Println("Incrementa en", x)

	//decremental
	x--
	fmt.Println("Decremental en", x)

	//Reto
	//calcular area del rectangulo, del trapecio, de un circulo

}
