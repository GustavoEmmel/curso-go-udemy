package main

import (
	"fmt"
)

// nao possuiu operador ternario em go
func obterResultado(nota float64) string {
	// return nota >= 6 ? "aprovado" : "reprovado"
	if nota >= 6 {
		return "aprovado"
	}
	return "reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}
