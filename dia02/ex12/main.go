// Exercicio 12 - Inferência de tipos: o Go infere o tipo de uma variável
// a partir do valor atribuído a ela.
package main

import "fmt"

func main() {
	v := 123i // change me!
	j := 456
	i := "olá"
	fmt.Printf("V is of type %T\n", v)
	fmt.Printf("J is of type %T\n", j)
	fmt.Printf("I is of type %T\n", i)
}
