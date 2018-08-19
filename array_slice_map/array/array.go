package main

import (
	"fmt"
)

func main() {
	// homogenea (mesmo tipo) e estatica (fixo)
	var notas [3]float64 //vai ser um array fixo de 3 posições e sempre com valores float 64
	fmt.Println(notas)
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas)
	// notas[3] = 10 nao vai ser compilado o programa por ser invalido o indice

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("media %.2f\n", media)
}
