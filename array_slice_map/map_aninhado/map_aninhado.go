package main

import (
	"fmt"
)

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gustavo":      880.55,
			"Guga Pereira": 11550.88,
		},
		"J": {
			"Joao": 3200.78,
		},
		"P": {
			"Pedro Rocha": 50000.45,
		},
	}

	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {
		//fmt.Println(letra, funcs)
		for nome, salario := range funcs {
			fmt.Println(letra, nome, salario)
		}
	}
}
