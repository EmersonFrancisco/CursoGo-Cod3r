package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro
	turboLigado bool
}

func main() {
	f1 := ferrari{}
	f1.nome = "F40"
	f1.velocidadeAtual = 250
	f1.turboLigado = true
	fmt.Println(f1)
	fmt.Printf("A Ferrari %s está a qual velocidade?\n", f1.nome)
	fmt.Printf("%dKM/h\n", f1.velocidadeAtual)
	fmt.Println("Ela está com o turbo ligado? ")
	if f1.turboLigado {
		fmt.Println("Sim... na BAGA")
	} else {
		fmt.Println("Não.")
	}
}
