package main

import (
	"fmt"
)

func multiplicacao(a, b int) int {
	return a * b
}

// usado em map reduce filter
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	fmt.Println(exec(multiplicacao, 2, 3))
}
