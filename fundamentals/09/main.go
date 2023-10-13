package main

import "fmt"

func main() {
	fmt.Printf("O valor total da soma é de %d", sum(1, 3, 4, 10, 80, 15))
}

// "..." dentro do escopo da função possibita receber, parâmetros infinitos.
func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		fmt.Printf("Somando %d ao valor total\n", numero)
		total += numero
	}
	return total
}
