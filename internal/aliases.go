package internal

import "fmt"

type AliasNumber int // declarando alias de inteiro

func Aliases() {
	var n AliasNumber

	n = 20

	fmt.Println(n)
}
