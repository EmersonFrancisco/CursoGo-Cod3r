package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(1, 5))

	subtrair := func(a, b int) int {
		return a - b
	}

	fmt.Println(subtrair(6, 7))
}
