package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1

	for i < 10 { // não tem While em Go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Print(" ", i)
	}

	fmt.Println("\n Usando IF dentro do FOR")

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("o Numero", i, "é Par")
		} else {
			fmt.Println("o Numero", i, "é Impar")
		}
	}
	tempo := 0
	for {
		//LAÇO INFINITO
		fmt.Println("Passou", tempo, "Segundo!!")
		tempo++
		time.Sleep(time.Second)
	}
}
