package internal

import "fmt"

// definindo struct do tiipo pessoa
type Pessoa struct {
	Nome  string
	Idade int
}

func Structs() {
	// criando uma instância de pessoa
	pessoa := Pessoa{
		Nome:  "Wesley",
		Idade: 45,
	}

	fmt.Println("Nome: ", pessoa.Nome, "Idade: ", pessoa.Idade)
}
