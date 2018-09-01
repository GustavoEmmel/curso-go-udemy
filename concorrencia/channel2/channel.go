package main

import (
	"fmt"
	"time"
)

// Channel (canal) - eh a forma de comunicacao entre as goroutines
// channel eh um tipo

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	fmt.Println("a")

	a, b := <-c, <-c // recebendo os dados do canal
	fmt.Println("b")
	fmt.Println(a, b)
	fmt.Println(<-c)

}
