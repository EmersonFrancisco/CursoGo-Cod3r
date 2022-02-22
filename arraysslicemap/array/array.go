package main

import "fmt"

func main() {
	//arrray é HOMOGENEA (MESMO TIPO) e estatica (o tamanho é sempre igual ao criado, não podendo ser acrescentado indice posteriormente)
	var nota [5]float64
	fmt.Println(nota)

	nota[0], nota[1], nota[2] = 7.4, 4.5, 5.9
	fmt.Println(nota)

	//fazer a media com todas notas dentro da array

	total := 0.0
	for i := 0; i < len(nota); i++ {
		total += nota[i]
	}
	media := total / float64(len(nota))
	fmt.Printf("A media é %2.f\n", media)
}
