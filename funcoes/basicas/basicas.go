package main

import "fmt"

//funcao sem nenhuma entrada, e nenhum retorno
func f1() {
	fmt.Println("F1")
}

//funcao com entrada dupla porem nenhum retorno
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s \n", p1, p2)
}

//funcao sem nenhuma entrada com retorno
func f3() string {
	return "F3"
}

//funcao com entrada dupla de forma abrevisada, e retorno
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

//funcao sem entrada e retorno duplo
func f5() (string, string) {
	return "retorno 1", "retorno 2"
}

func main() {
	f1()
	f2("teste", "teste2")
	r3, r4 := f3(), f4("teste f4", "teste2 f4")
	fmt.Println(r3)
	fmt.Println(r4)
	r51, r52 := f5() // pode-se ignorar o retorno de 1 dos paramentros usando um _ na definição da variavel
	fmt.Println("F5", r51, " e ", r52)
}
