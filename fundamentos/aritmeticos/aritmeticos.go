package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Modulo =>", a%b)

	//bitwise
	fmt.Println("E =>", a&b)   // 11 & 10 = 10
	fmt.Println("Ou =>", a|b)  // 11 | 10 = 11
	fmt.Println("Xor =>", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	// outras operacoes usando math
	fmt.Println("Maior =>", math.Max(float64(a), float64(b))) // maior valor entre a e b e foi convertido de int para float
	fmt.Println("Menor =>", math.Min(c, d))                   // menor valor entre c e d
	fmt.Println("Exponenciação =>", math.Pow(c, d))           // c elevado por d

}
