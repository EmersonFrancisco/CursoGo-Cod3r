package main

import "fmt"

func fatorial(num int) (int, error) {
	switch {
	case num < 0:
		return -1, fmt.Errorf("Número inválido: %d", num)
	case num == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(num - 1)
		return num * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-4)

	if err != nil {
		fmt.Println(err)
	}
}
