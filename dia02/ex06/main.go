// Exercicio 6 - Retorno limpo (naked return) e retorno nomeado para simplificar
// a declaração de variáveis.
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
