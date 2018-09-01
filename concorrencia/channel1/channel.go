package main

import (
	"fmt"
)

func main() {
	// tipo de canal (valores inteiros), buffer
	ch := make(chan int, 1)

	ch <- 1 // estou enviando para o canal o valor 1 (escrita)
	<-ch    // recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
