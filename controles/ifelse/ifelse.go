package main

import "fmt"

func main() {
	imprimirresultado(7.9)
	imprimirresultado(5.9)
}

func imprimirresultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}

}
