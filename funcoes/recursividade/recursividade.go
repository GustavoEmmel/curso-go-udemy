package main

import (
	"fmt"
)

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("numero invalido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	retsultado, _ := fatorial(5)
	fmt.Println(retsultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}

	// uma solução melhor seria inteiro sem sinais
}
