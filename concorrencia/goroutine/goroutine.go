package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteracao) %d\n", pessoa, texto, i+1)
	}
}

func main() {
	// estamos usando a funcao de maneira serial
	// fale("Jhon", "Fale comigo", 3)
	// fale("Doe", "So posso falar depois de voce", 1)

	// go fale("Jhon", "Hey Ho", 500)
	// go fale("Doe", "Lets Go", 500)

	// time.Sleep(time.Second * 5)
	// fmt.Println("Fim")

	go fale("Maria", "entendi", 10)
	fale("joao", "parabens", 5)
}
