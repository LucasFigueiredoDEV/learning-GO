package main

import "fmt"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Lucas"
	e float64 = 1.2
	f ID      = 1
) //crianção de variável global

func main() {
	fmt.Printf("O tipo da variável E é %T", e)
}

// %T retorna o tipo
// %v retorna o valor
