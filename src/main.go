package main

import "fmt"

func normalFunction(mensaje string) {
	fmt.Println(mensaje)
}

func tripleArgunment(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func areaRec(a, b int) int {
	return a * b
}

func areaTrap(a, b, h int) int {
	return ((a + b) * h) / 2
}

func areaCirc(p, r float64) float64 {
	return p * r * r
}

func main() {
	//actividad con funciones
	arearectangulo := areaRec(15, 20)
	fmt.Println("el area del rectangulo en función es: ", arearectangulo)

	areatrap := areaTrap(20, 10, 10)
	fmt.Println("el area del trapecio en función es:", areatrap)

	areacirc := areaCirc(3.141516, 10)
	fmt.Println("el area del circulo en función es: ", areacirc)

	normalFunction("Hola mundo")
	tripleArgunment(40, 20, "doble hola mundo")

	value := returnValue(2)
	fmt.Println("el valor es", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("los numeros son", value1, value2)

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
	const BaseRectangulo = 20
	const AlturaRectangulo = 10
	AreaRectangulo := BaseRectangulo * AlturaRectangulo
	fmt.Println("El area del rectangulo es: ", AreaRectangulo)

	const LongA = 10
	const LongB = 15
	const AlturaTrapecio = 20
	Sumalados := LongA + LongB
	AreaTrapecio := (Sumalados * AlturaTrapecio) / 2
	fmt.Println("El area del trapecio es: ", AreaTrapecio)

	const radio = 5
	var AreaCirulo float64 = pi2 * radio * radio
	fmt.Println("El area del circulo es: ", AreaCirulo)

	// paquete fmt
	// Decalración de variables
	JelouMessage := "Jelou"
	var WorldMessage string = "World"

	// Println - salto de linea automatica
	fmt.Println(JelouMessage, WorldMessage)
	fmt.Println(JelouMessage, WorldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos) // Contra slash con alt+92 \
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos) // si no se que tipo de dato voy a imprimir

	//Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Println(message)

	// Tipo de Dato
	fmt.Printf("JelouMessage: %T\n", JelouMessage)
	fmt.Printf("radio: %T\n", radio)

	// Uso de funciones

}
