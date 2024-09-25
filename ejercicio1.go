package main

import "fmt"

func imprimir() {
	fmt.Println("Texto desde la funcion imprimir")
}

func hola_string(s string) string {
	return "Hola " + s

}

func sumar(a int, b int) int {
	return a + b
}

func comparar_bool() {
	var a bool
	b := true
	a = false
	if a == b {
		fmt.Println("A y B son iguales")
	} else {
		fmt.Println("A y B son difirentes")
	}
}

// ejercicio 15
func variables() {
	var Nombre, Apellido string = "Katherine", "Rincon"
	fmt.Println("nombre: ", Nombre, "\n Apellido: ", Apellido)

}
func Datos() {

	var (
		Nombre     string = "Darling"
		Edad       int    = 30
		Pensionado bool   = false
	)
	fmt.Println("Nombre: ", Nombre)
	fmt.Println("Edad: ", Edad)
	fmt.Println("Pensionado: ", Pensionado)
}

// Ejercicio 16

func Informacion() {
	var Nombre string
	var Edad int
	var Peso float64
	var Estudiante bool
	fmt.Println("Nombre: ", Nombre)
	fmt.Println("Edad: ", Edad)
	fmt.Println("Peso: ", Peso)
	fmt.Println("Estudiante: ", Estudiante)

}
// Ejercicio 17
func Operador() {
	Nombre := "David Suarez"
	Edad := 30
	Peso := 75
	Estudiante := false
	fmt.Println("Nombre: ", Nombre)
	fmt.Println("Edad: ", Edad)
	fmt.Println("Peso: ", Peso)
	fmt.Println("Estudiante: ", Estudiante)
}
// Ejercicio 18
var var1 ="Este es el nivel 1"

func Nivel() {
	var var2 ="Este es el nivel 2"
	{
		var var3 = "Este es el nivel 3"
		fmt.Println(var3)
	}
	fmt.Println(var1)
	fmt.Println(var2)
}

// Ejercicio 19
func Vehiculos(){

	vehiculo1 := "rojo"
	fmt.Println("El vehiculo1 es",vehiculo1)

	vehiculo2 := vehiculo1
	fmt.Println("El vehiculo2 es", vehiculo2)

	vehiculo3 := &vehiculo1
	fmt.Println("El vehiculo es", *vehiculo3)

	vehiculo1 = "gris"

	fmt.Println("El vehiculo1 es", vehiculo1, &vehiculo1)
	fmt.Println("El vehiculo2 es", vehiculo2, &vehiculo2)
	fmt.Println("El vehiculo3 es", *vehiculo3, vehiculo3)
}	
func main() {
	// fmt.Println("imprimir desdes la funcion main")
	// imprimir()
	//fmt.Println(hola_string("Kathe"))
	//fmt.Println(sumar(3, 5))
	// comparar_bool()
	// variables()
	// Datos()
	// Informacion()
	// Operador()
	// Nivel()
       Vehiculos() 
}

