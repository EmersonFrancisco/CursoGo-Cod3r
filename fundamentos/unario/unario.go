package main

import "fmt"

func main() {
	x := 1
	y := 2

	//go ten penas operadores postfix
	x++ // seria a mesma coisa que x = x + 1
	fmt.Println(x)
	y-- // seria a mesma coisa que y = y - 1
	fmt.Println(y)

	//fmt.Println(x == y--) não funciona uma postfix dentro de uma comparaçao, fazer um de cada vez
}
