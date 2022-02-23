package main

import "fmt"

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4 // linha de bloqueio pq encheu os 3 buffer
	ch <- 5
	fmt.Println("Execultou")
	ch <- 6
	ch <- 7
	ch <- 8
	ch <- 9
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)
	fmt.Println(<-ch)

}
