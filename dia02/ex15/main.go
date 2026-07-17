// Exercício 15 — divmod

// Crie a função divmod(a, b int). Ela retorna 3 valores: quociente, resto e erro.
// Caminho feliz: chame divmod(17, 5). Deve imprimir 3 2.
// Caminho triste: chame divmod(10, 0). O programa não pode quebrar. Deve imprimir uma mensagem de erro sua, criada com fmt.Errorf.
// Trate o erro com if err != nil. Proibido usar panic.

package main

import "fmt"

func divmod(a, b int) (quociente, resto int, err error) {
	if b == 0 {
		err = fmt.Errorf("não é possível dividir por zero")
		return
	}

	quociente = a / b
	resto = a % b
	return

}

func main() {
	q, r, err := divmod(17, 5)
	// q, r, err := divmod(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(q, r)
	}
}
