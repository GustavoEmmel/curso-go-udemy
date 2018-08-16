package main

import (
	"fmt"
)

func main() {
	x := 1
	y := 2

	//go somente possui apenas pos fixada
	// ++x nao existe (pre fixada)
	x++ // x += 1 ou x = x + 1
	fmt.Println(x)

	y--
	fmt.Println(y)

	//fmt.Println(x == y--) nao pode
}
