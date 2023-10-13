package main

import "fmt"

var first_number = 8

var second_number = 3

var result = sum(first_number, second_number)

func main() {
	fmt.Println(sum_error(1, 2))

	fmt.Printf("O valor da soma entre os números %d e %d é de %d", first_number, second_number, result)
}

func sum(a int, b int) int {
	return a + b
}

func sum_error(a int, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
} // Primeiro parêntese representa os parâmetros que serão passados e seus tipos
// Segundo parêntese representa o tipo que será retornado
