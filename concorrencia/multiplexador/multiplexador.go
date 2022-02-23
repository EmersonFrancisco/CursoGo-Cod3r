package main

import (
	"emerson/urls"
	"fmt"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}

}

//multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		urls.Titulo("https://www.amazon.com", "https://www.youtube.com"),
		urls.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
