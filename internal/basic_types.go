package internal

import "fmt"

func TiposBasicos() {
	var idade int = 30             // inteiro
	var numero int64 = 10000000000 // inteiro 64 bits

	var altura float32 = 1.65             // ponto flutuante 32 bits
	var distancia float64 = 256.123456789 // ponto flutuante 64 bits

	var estaChovendo bool = true // booleano

	var nome string = "Wesley" // string

	valorDinamico := "teste" // maneira alternativa de criar variáveis e inserir seus valores sem a necessidade de explicitar o tipo

	fmt.Println("idade: ", idade, "número: ", numero, "altura: ", altura, "distancia: ", distancia, "esta chovendo? ", estaChovendo, "nome: ", nome, "variável dinamica: ", valorDinamico)
}
