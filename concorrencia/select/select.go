package main

import (
	"emerson/urls"
	"fmt"
	"time"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := urls.Titulo(url1)
	c2 := urls.Titulo(url2)
	c3 := urls.Titulo(url3)

	// estrutura de controle especifica para concorrencia
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam"
		//	default:
		//		return "Sem resposta ainda!!!"
	}

}

func main() {
	campeao := oMaisRapido(
		"https://www.amazon.com",
		"https://www.youtube.com",
		"https://www.google.com",
	)
	fmt.Println(campeao)
}
