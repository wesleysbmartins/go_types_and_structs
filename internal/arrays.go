package internal

import "fmt"

func Arrays() {
	var numeros [3]int // declarando array de inteiros do tamanho 3

	// atribuição de valores
	numeros[0] = 1
	numeros[1] = 2
	numeros[2] = 3

	fmt.Println("Números: ", numeros)
}
