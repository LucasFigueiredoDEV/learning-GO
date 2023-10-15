package main

import "fmt"

func main() {
	lucas := Cliente{
		Nome:  "Lucas",
		Idade: 20,
		Ativo: true,
	}
	lucas.Cidade = "Recife"
	lucas.Ativo = false         // É possível mudar os valores que estão inseridos na Struct
	lucas.Estado = "Pernambuco" // Tem a possibilidade de criar composições
	lucas.Desativar()
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", lucas.Nome, lucas.Idade, lucas.Ativo)
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}
