package internal

import "fmt"

// struct
type Animal struct {
	Nome  string
	Idade int
	Tipo  string
}

// interface com contratos de métodos a serem implementados
type IAnimal interface {
	Speak()
}

// método
func (a *Animal) Speak() {
	if a.Tipo == "Cão" {
		fmt.Println("AAUUUU!")
	} else if a.Tipo == "Gato" {
		fmt.Println("MIAAAU!")
	}
}
