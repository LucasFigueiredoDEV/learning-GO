package main

import "fmt"

func main() {
	salarios := map[string]int{"Lucas": 3000, "João": 2500, "Maria": 5000}

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é de R$%d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é de R$%d\n", salario)
	} // O "_" no for é utilizado para que o valor do dict/map seja ignorado
}
