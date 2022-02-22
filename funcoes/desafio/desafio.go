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
	resultado, err := fatorial(-5)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}

	// ou usar o uint assim não aceitaria numero negativo na função, não precisando de tal verificação, porem caso o usuario
	// final colocar esse valor, não iria retornar erro, e não saberia o motivo!!!!!!!
	// a não ser que coloque uma verifacação logo após a inserção do valor
}
