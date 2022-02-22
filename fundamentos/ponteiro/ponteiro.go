package main

import "fmt"

func main() {
	i := 1

	//go não tem aritimeticas de ponteiros
	var p *int = nil
	p = &i
	*p++
	i++

	//p++ não funciona..poois go não tem aritimeticas de ponteiros

	fmt.Println(p, *p, i, &i)
}
