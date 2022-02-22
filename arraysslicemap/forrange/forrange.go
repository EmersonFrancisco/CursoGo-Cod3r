package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilator conta o tamanho da array

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1 /*o +1 é apenas para somar 1 numero ai indice, tornando 0 em 1 para melhor aparencia*/, numero)
	}
	for _, numero := range numeros { // o _ anula a variavel de indice
		fmt.Println(numero) //dessa forma você não precisa obrigatoriamente usar o numero do indece, podendo apenas pegar o valor!!
	}
}
