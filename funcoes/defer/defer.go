package main

import "fmt"

func obtorValorAprovado(numero int) int {
	defer fmt.Println("Fim!!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo")
	return numero
}

func main() {
	fmt.Println(obtorValorAprovado(6000))
	fmt.Println(obtorValorAprovado(3000))
}
