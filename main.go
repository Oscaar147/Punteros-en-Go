package main

import "fmt"

func intercambia(a *int, b *int) { // se recibe un puntero
	aux := 0
	aux = *a
	*a = *b
	*b = aux

}

func main() {
	a := 0
	b := 0

	fmt.Println("Ingresa el entero A")
	fmt.Scan(&a)
	fmt.Println("Ingresa el entero B")
	fmt.Scan(&b)

	intercambia(&a, &b) // con & se obtiene la dirección de memoria de una variable

	fmt.Println("El numero A es: ", a)
	fmt.Println("El numero B es: ", b)
}
