package main

import "fmt"

func main() {
	double_total := func() int {
		return sum(1, 3, 4, 10, 80, 15) * 2
	}()

	fmt.Printf("O valor total da soma multiplicado por 2 é de %d", double_total)
}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		fmt.Printf("Somando %d ao valor total\n", number)
		total += number
	}
	return total
}
