// Exercício 17 — overflow: variável vs constante

// Programa A: var x int8 = 127, depois x = x + 1, depois imprima x.
// Programa B: var y int8 = 127 + 1, depois imprima y.
// Antes de rodar, preveja cada um: compila ou não? Se compila, imprime o quê?
// Rode os dois separados. Um se comporta diferente do outro.
// Me explique o porquê da diferença. Dica: no B, o 127 + 1 é uma constante. Lembra da aula do 2¹⁰⁰.

package main

import "fmt"

func main() {
	var x int8 = 127
	x = x + 1
	fmt.Println(x) // Variável só estoura em tempo de execução, então o Go permite o overflow e imprime -128

	// var y int8 = 127 + 1
	// fmt.Println(y) // Constante estoura em tempo de compilação, então o Go não permite o overflow e dá erro de compilação
}
