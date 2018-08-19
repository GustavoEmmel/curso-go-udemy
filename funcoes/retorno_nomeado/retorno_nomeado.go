package main

import (
	"fmt"
)

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1

	return // retorno limpo pois ja foi atribuido o valor de retorno para as variaveis
}

func main() {
	seg, pri := trocar(1, 2)
	fmt.Println(pri, seg)
}
