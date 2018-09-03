package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operacao bloqueante
	fmt.Println("so depois que for lido")
}

func main() {
	c := make(chan int) // canal sem buffer (o primeiro dado do canal ja ira bloquear)

	go rotina(c)

	fmt.Println(<-c) // operacao bloqueante
	fmt.Println("foi lido")
	fmt.Println(<-c) //deadlock
	fmt.Println("fim")
}
