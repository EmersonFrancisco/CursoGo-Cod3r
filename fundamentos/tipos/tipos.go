package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 ..16 ... 32.... 64
	var b byte = 255
	fmt.Println("o Byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("o valor máximo do int é", i1)
	fmt.Println("o tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela unicode (int32)
	fmt.Println("o Rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numero reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // padrão é float64

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// String
	s1 := "Testante a var" // go não permite delimitar string com aspas simples > '' <
	fmt.Println("o tipo da s1 é", reflect.TypeOf(s1))
	fmt.Println("o tamanho da string é", len(s1)) // saber quantidade de caracteres da string

	// String com multiplas linhas
	s2 := `Olá
	meu
	nome
	é Emerson` // para usar multiplas linhas se coloca uma CRASE, > ` ` <
	fmt.Println("o tamanho da string é", len(s2))
	fmt.Println(s2)

	// char??
	// não tem

}
