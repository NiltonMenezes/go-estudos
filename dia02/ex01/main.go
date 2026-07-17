// Exercicio 1 - Entendendo pacotes e as diferenças entre Println, Printf e Print,
// além da posição de cada pacote em disco.
package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Número aleatório:", rand.Intn(10))
	fmt.Printf("Agora você tem %g problemas. \n", math.Sqrt(7))
	fmt.Println(math.Pi)
}
