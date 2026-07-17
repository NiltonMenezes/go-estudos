// Exercicio 2 - Entendendo os valores padrão de cada tipo de dado e como eles são
// representados na saída padrão. Além disso, a diferença entre %v e %q na formatação de strings.
package main

import "fmt"

var i int
var s string
var b bool
var f float64
var sl []int

func main() {
	fmt.Printf("%v | %v | %v | %v | %v\n", i, s, b, f, sl)
	fmt.Printf("%q\n", s) // repare a diferença do %q para o %v na string
}
