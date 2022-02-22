package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ") //usando o strings.split, cria uma slice com as palavras como por exemplo forem
	//separadas por um espaço que foi informado em " "
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{
		"Emerson",
		"Francisco",
	}
	fmt.Println(p1.getNomeCompleto())
	p1.setNomeCompleto("Samara Matos")
	fmt.Println(p1.getNomeCompleto())
}
