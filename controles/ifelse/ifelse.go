package main

import (
	"fmt"
)

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("aprovado com nota", nota)
	} else {
		fmt.Println("reprovado com nota", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.3)
}
