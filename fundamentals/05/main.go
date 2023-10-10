package main

import "fmt"

func main() {
	var meuArray [3]int // Definindo tamanho do array
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 10

	fmt.Println(meuArray[0])               // printando o primeiro valor do array
	fmt.Println(meuArray[1])               // printando o segundo valor do array
	fmt.Println(meuArray[2])               // printando o terceiro valor do array
	fmt.Println(meuArray[len(meuArray)-1]) // printando o último valor do array/lista

	for i, v := range meuArray {
		fmt.Printf("O valor do indice %d é de %d\n", i, v) // \n para quebrar uma linha
	} // Percorrendo array com o for
}
