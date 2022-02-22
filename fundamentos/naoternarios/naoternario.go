package main

import "fmt"

//não tem operador ternario
func obterresultado(nota float64) string {
	//return nota >= 6 ? "aprovador" : "reprovado"
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterresultado(6.7))

}
