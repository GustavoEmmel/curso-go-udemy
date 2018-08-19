package main

import (
	"fmt"
)

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //compilador vai contar a quantidade de itens no array

	// equivalente ao foreach
	// primeiro elemento eh a chave
	// segundo elemento eh o valor

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
