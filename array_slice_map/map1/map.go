package main

import (
	"fmt"
)

func main() {
	//var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678900] = "jon"
	aprovados[12345678911] = "doe"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678900])
	delete(aprovados, 12345678911)
	fmt.Println(aprovados[12345678911]) //nao vai imprimir nada
}
