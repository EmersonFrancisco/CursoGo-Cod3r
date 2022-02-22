package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) ingerido opelo compilador

	// forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)
	fmt.Println("A Area da circunferencia é", area)

	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	//criando mais de uma variavel na mesma linha
	var e, f bool = true, false

	fmt.Println(e, f)

	//criando mais de uma variavel de tipos diferente e resumida na mesma linha
	g, h, i := 2, false, "caramba que top!!!"

	fmt.Println(g, h, i)
}
