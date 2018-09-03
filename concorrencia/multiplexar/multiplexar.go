package main

import (
	"fmt"

	"github.com/GustavoEmmel/html"
)

func encaminhar(origin <-chan string, destino chan string) {
	for {
		destino <- <-origin
	}
}

// multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo(
			"https://www.google.com/",
			"https://bematech.bemacash.com.br/",
		),
		html.Titulo(
			"https://www.amazon.com.br/",
			"https://www.cod3r.com.br/",
		),
	)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
