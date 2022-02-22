package main

import "fmt"

func main() {
	fmt.Print("mesma")
	fmt.Print(" linha")

	fmt.Println(" Nova")
	fmt.Println(" linha")
	x := 3.141316

	//fmt.Println("0 valor de x é" + x) não funciona
	// pode se fazer a parte abaixo para usar o + VAR
	// porem é mais pratico usar a ,
	xs := fmt.Sprint(x)
	//Sprint retorna a informação da variavel como uma String
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)

	//usando o argumento %.3f você define que a var
	//a serguir mencionada ira exibir apenas 3 casas
	// após a virgula
	// %f float (numero fracionado)
	// %s String (texto)
	// %d Inteiro (numero não fracionado)
	// %t Boolean (false ou true)
	// \n pular linha dentro da mesma string
	// %v inteiro, float sem definir qt de casa, boolean, string

	fmt.Printf("O valor de x é %.3f", x)

	//PrintF - usar o mecanismo %...
	//Print - printa a informação
	//Println - printa a informação e pula linha
	a := 1
	b := 1.999
	c := "texto aqui"
	d := false
	fmt.Printf("\n%d \n%f \n%.2f \n%t \n%s", a, b, b, d, c)
	fmt.Printf("\n%v \n%v \n%v", a, c, b)
}
