package internal

import "fmt"

func If() {
	n := 0
	b := false
	l := []string{"x", "y"}

	// validação de inteiros
	if n > 0 {
		fmt.Println("N é maior que 0!")
	} else {
		fmt.Println("N não é maior que 0!")
	}

	n = 30

	if n > 0 {
		fmt.Println("N é maior que 0!")
	} else {
		fmt.Println("N não é maior que 0!")
	}

	// validação de bool
	if b {
		fmt.Println("B é verdadeiro!")
	} else {
		fmt.Println("B é falso!")
	}

	b = true

	if !b {
		fmt.Println("B é falso!")
	} else {
		fmt.Println("B é verdadeiro!")
	}

	// validação do slice
	if len(l) > 0 {
		fmt.Println("Slice: ", l)
	} else {
		fmt.Println("O Slice esta vazio!")
	}

	emptyL := []string{}

	if len(emptyL) > 0 {
		fmt.Println("Slice: ", emptyL)
	} else {
		fmt.Println("O Slice esta vazio!")
	}
}
