package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notafinal := int(nota)
	fmt.Println(notafinal)

	//cuidado...
	//fmt.Println("teste" + string(97)) dessa forma ira imprimir o caracterer que se encontra na casa 97 no dicionario de caracteres

	//int para string
	fmt.Println("teste" + strconv.Itoa(int(nota)))

	//string para int
	num, _ := strconv.Atoi("123")
	//num recebe o int, e _ retorna um erro caso a string não seja inteiro
	fmt.Println(num - 122)

	//String para Boolean
	b, _ := strconv.ParseBool("false")
	if b {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("Falso")
	}
}
