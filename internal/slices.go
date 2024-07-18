package internal

import "fmt"

func Slices() {
	frutas := []string{"Maçã", "Banana", "Morango"} // declarando slice

	frutas = append(frutas, "Pera") // adicionando um elemento ao slice

	// iterando slice
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}
}
