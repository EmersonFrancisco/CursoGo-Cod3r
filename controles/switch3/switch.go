package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Numero Inteiro"
	case float32, float64:
		return "Numero Real"
	case string:
		return "Texto String"
	case func():
		return "Uma Função"
	default:
		return "Não sei que tipo é esse!!"
	}
}
func main() {
	fmt.Println(tipo(1.8))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Ola"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
