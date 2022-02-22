package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo
}

func main() {
	segundo, primeiro := trocar(2, 5)
	fmt.Println(segundo, primeiro)
}
