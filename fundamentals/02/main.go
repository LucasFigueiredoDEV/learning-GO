package main

const a = "Hello, world!"

var b bool = true
var c int = 10
var d string = "Lucas"

func main() {
	e := 1.2 // Com ":=" é possível criar variáveis sem especificar o tipo, pois a própria linguagem consegue saber qual o tipo (Só é válido declarar váriavel dessa maneira dentro das funções)
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
}
