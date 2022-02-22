package main

import "fmt"

func main() {
	tv50, tv32, sorvete := compras(true, true)

	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete %t, Salvadel: %t", tv50, tv32, sorvete, !sorvete)

}

func compras(trab1, trab2 bool) (bool, bool, bool) {
	compraTV50 := trab1 && trab2
	compraTV32 := trab1 != trab2 // ou exclusivo
	compraSorvete := trab1 || trab2

	return compraTV50, compraTV32, compraSorvete
}
