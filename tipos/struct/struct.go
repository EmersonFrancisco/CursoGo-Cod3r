package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.50,
		desconto: 0.05,
	}
	produto2 := produto{"Notebook", 1900.99, 0.10}
	produto3 := produto{"Monitor", 580.99, 0.10}

	var listaProduto [3]produto

	listaProduto[0] = produto1
	listaProduto[1] = produto2
	listaProduto[2] = produto3
	fmt.Println("Lista de produtos")
	for i := 0; i < len(listaProduto); i++ {
		fmt.Println("Produto ", i+1, ":")
		fmt.Printf("Nome: %s\n", listaProduto[i].nome)
		fmt.Printf("Preço: R$%.2f\n", listaProduto[i].preco)
		fmt.Printf("Preço Avista: R$%.2f\n\n", listaProduto[i].precoComDesconto())
	}

}
