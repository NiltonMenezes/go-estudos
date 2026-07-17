// Exercicio 4 - Funções com múltiplos parâmetros e como simplificar a declaração
// de tipos iguais em parâmetros consecutivos.
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(20, 25))
}
