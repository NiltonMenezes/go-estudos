// Exercício 16 — conversão de float para inteiro

// Escreva este código: f := 3.999 e imprima uint(f). Depois: g := -3.7 e imprima int(g).
// Antes de rodar, escreva num comentário o que você acha que sai em cada um.
// Rode. Depois me responda em uma frase: o Go arredonda ou corta os decimais?

package main

import "fmt"

func main() {
	f := 3.999
	fmt.Println(uint(f)) // Acho que vai sair um 3. O unit é inteiro sem sinal, não pode ser negativo, então o Go vai cortar os decimais e imprimir 3
	g := -3.7
	fmt.Println(int(g)) // Acho que vai sair um -3. É inteiro com sinal, então o Go vai cortar os decimais e imprimir -3

	// O Go corta os decimais, não arredonda
	// int8 tem 8 bits. 1 bit vai para o sinal. Sobram 7 para o valor → 2⁷ = 128 combinações de cada lado. Faixa: −128 até 127. O máximo é 2⁷ − 1 porque o zero ocupa uma vaga do lado positivo.
}
