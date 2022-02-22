package main

import "fmt"

func main() {
	funceSalario := map[string]float64{
		"Samara Matos":      2400.00,
		"Emerson Francisco": 1980.00,
		"Fernando Junior":   1200.00,
	}
	fmt.Println("Lista de Funcionarios / Salarios")
	for nome, salario := range funceSalario {
		fmt.Printf("%s - R$%.2f\n", nome, salario)
	}
	funceSalario["Ana Julia Matos"] = 1250.40

	delete(funceSalario, "Samara Matos")

	fmt.Println("\nNova Lista de Funcionarios / Salarios")
	for nome, salario := range funceSalario {
		fmt.Printf("%s - R$%.2f\n", nome, salario)
	}

}
