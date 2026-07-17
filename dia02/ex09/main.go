// Exercicio 9 - Operador de declaração curta (:=) para simplificar
// a declaração de variáveis.
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
