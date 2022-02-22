package main

import (
	"fmt"
)

type item struct {
	itemID     int
	quantidade int
	preco      float64
}

type pedido struct {
	pedidoID int
	itens    []item
}

func (p pedido) precoTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}
	return total
}

func main() {
	pedido := pedido{
		pedidoID: 1,
		itens: []item{
			{1, 2, 9.50},
			{2, 1, 30.00},
			{3, 10, 1.50},
		},
	}
	fmt.Printf("O valor total do pedido foi: R$%.2f", pedido.precoTotal())
}
