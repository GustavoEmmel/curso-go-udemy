package main

import (
	"fmt"
)

func main() {
	funcsESalarios := map[string]float64{
		"Jon":    11333.45,
		"Doe":    8000.44,
		"Junior": 1200.0,
	}

	funcsESalarios["Chapolin"] = 5550.00
	// nao tem erro
	delete(funcsESalarios, "inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
