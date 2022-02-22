package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados como abaixo:::
	aprovados := make(map[int]string)

	aprovados[12345678989] = "Maria"
	aprovados[98765432112] = "João"
	aprovados[98712345609] = "Henrique"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
	fmt.Println("\nvocê vai deletar do registro o aprovado -> ", aprovados[12345678989])
	delete(aprovados, 12345678989)

	fmt.Println("\nNova Lista de Aprovados:")
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

}
