package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) <= inicio && float64(n) >= fim
}

func notaParaConceito(n nota) string {
	switch {
	case n.entre(10.0, 9.0):
		return "A"
	case n.entre(8.99, 7.00):
		return "B"
	case n.entre(7.99, 5.00):
		return "C"
	case n.entre(4.99, 3.00):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(4.6))
}
