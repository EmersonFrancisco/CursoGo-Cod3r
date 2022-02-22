package main

import (
	"area"
	"fmt"
)

func main() {
	fmt.Printf("a Area do circulo é : %.2fr²\n", area.Circ(6.0))
	fmt.Printf("a Area do Retangulo é: %.2fm²\n", area.Rect(2, 8))
	fmt.Printf("a Area do Quadrado é: %.2fm²\n", area.Quad(3))
	fmt.Printf("a Area do Cubo é: %.2fm³\n", area.Cubo(3, 3, 4))
	fmt.Printf("a Area do Triangulo Equilátero é: %.2fm²", area.TrianguloEq(3, 3))
}
