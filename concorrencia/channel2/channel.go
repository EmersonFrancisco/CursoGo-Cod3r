package main

import (
	"fmt"
	"time"
)

func multi(num int, ch chan int) {
	time.Sleep(time.Second)
	ch <- 2 * num //enviando dados para o canal
	time.Sleep(time.Second)
	ch <- 3 * num // enviando novamente dados para o canal
	time.Sleep(3 * time.Second)
	ch <- 4 * num // enviando novamente dados para o canal
}

func main() {
	c := make(chan int)
	go multi(2, c)

	a, b := <-c, <-c // recebendo os dados do canal
	fmt.Println(a, b)
	fmt.Println(<-c)
}
