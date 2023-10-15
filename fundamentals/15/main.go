package main

func main() {
	// Memória -> Endereço -> Valor

	// variável -> Ponteiro que tem um endereço na memória -> valor
	a := 10
	var ponteiro *int = &a // *int serve para ver o endereçamento na memória
	*ponteiro = 20         //Trocando o valor da variável pelo endereçamento
	println(a)

	b := &a
	*b = 30
	println(*b)
	println(a)
}
