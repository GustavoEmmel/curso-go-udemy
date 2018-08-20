package main

import (
	"fmt"
)

func inc1(n int) {
	n++
}

// ponteiro eh representado por um *

func inc2(n *int) {
	// revisao: * eh usado para desreferenciar, ou seja
	// ter acesso ao valor no qual o ponteiro aponta

	*n++
}

func main() {
	n := 1
	inc1(n)

	fmt.Println(n)

	// revisao & eh usado para o endereco de uma variavel
	inc2(&n) //por referencia
	fmt.Println(n)
}
