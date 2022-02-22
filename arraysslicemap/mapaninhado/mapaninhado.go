package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"S": {
			"Samara Matos": 2560.09,
			"Souza Junior": 1200.60,
		},
		"E": {
			"Emerson Francisco": 2500.50,
			"Ernesto Silva":     1400.60,
		},
		"J": {
			"João Jorge":  1540.80,
			"Julio Cesar": 1500.50,
		},
	}
	fmt.Println("Lista de Funcionarios Organizados por Letra")
	for letra, funcio := range funcsPorLetra {
		fmt.Printf("Letra %s:\n", letra)
		for nome, salario := range funcio {
			fmt.Printf("%s - R$%.2f\n", nome, salario)
		}
	}
	deletar := "J"
	fmt.Println("\nVocê ira deletar todos funcionarios com a letra", deletar)
	delete(funcsPorLetra, deletar)

	fmt.Println("\nNova Lista de Funcionarios Organizados por Letra")
	for letra, funcio := range funcsPorLetra {
		fmt.Printf("Letra %s:\n", letra)
		for nome, salario := range funcio {
			fmt.Printf("%s - R$%.2f\n", nome, salario)
		}
	}
}
