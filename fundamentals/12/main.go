package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

func main() {
	lucas := Cliente{
		Nome:  "Lucas",
		Idade: 20,
		Ativo: true,
	}
	lucas.Cidade = "Recife"
	lucas.Ativo = false // É possível mudar os valores que estão inseridos na Struct
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", lucas.Nome, lucas.Idade, lucas.Ativo)
}
