/*
package main

import "fmt"

	func suna(a, b int) int {
		return a + b
	}

	func main() {
		nombre := "facundo"
		edad := 20
		fmt.Println("hola, ", nombre, "! tenes ", edad, " anios\n")
		resultado := suna(5, edad)
		fmt.Println("la suma es: ", resultado)
	}
*/
/*package main

import "fmt"

func calculo() (int, bool) {
	return 10, true
}
func main() {
	a, b := calculo()
	fmt.Println(a, b)
}*/
/*func main() {
	var input string
	fmt.Println("ingrese un valor")
	fmt.Scan(&input)
	fmt.Println("valor ingresado ", input)
}*/
package main

import "fmt"
import "math"

func hipotenusa(lado1, lado2 float64) float64 {
	return math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))
}
func area(lado1, lado2 float64) float64 {
	return (lado1 * lado2) / 2
}
func perimetro(lado1, lado2, hipotenusa float64) float64 {
	return lado1 + lado2 + hipotenusa
}
func main() {
	fmt.Println("ingrese los lados del triangulo")
	var lado1 float64
	var lado2 float64
	fmt.Println("ingrese el lado1")
	fmt.Scan(&lado1)
	fmt.Println("ingrese el lado2")
	fmt.Scan(&lado2)
	Hipotenusa := hipotenusa(lado1, lado2)
	fmt.Println("la hipotenusa es ", Hipotenusa)
	Area := area(lado1, lado2)
	fmt.Println("el area es ", Area)
	Perimetro := perimetro(lado1, lado2, Hipotenusa)
	fmt.Println("el perimetro es ", Perimetro)
	fmt.Printf("la hipotenusa es:%.2f\n", Hipotenusa)
	fmt.Printf("el area es:%.2f\n", Area)
	fmt.Printf("el perimetro es:%.2f\n", Perimetro)
}
