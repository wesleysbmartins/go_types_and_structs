package main

import (
	"fmt"
	"go_types_and_structs/internal"
)

func main() {
	internal.TiposBasicos()
	internal.Constants()
	internal.Aliases()
	internal.Arrays()
	internal.Slices()
	internal.Maps()
	internal.Structs()

	animal := internal.Animal{
		Nome:  "Half",
		Idade: 8,
		Tipo:  "Cão",
	}

	fmt.Println("ANIMAL: ", animal)

	animal.Speak()
}
