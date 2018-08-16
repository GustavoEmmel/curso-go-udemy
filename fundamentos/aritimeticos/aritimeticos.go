package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("soma => ", a+b)
	fmt.Println("subtracao => ", a-b)
	fmt.Println("divisao => ", a/b)
	fmt.Println("multiplicacao => ", a*b)
	fmt.Println("modulo (resto) => ", a%b)

	//Bitwise
	fmt.Println("E => ", a&b)   // 11 & 10 = 10
	fmt.Println("Ou => ", a|b)  // 11 | 10 = 11
	fmt.Println("Xor => ", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	fmt.Println("maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("menor =>", math.Min(c, d))
	fmt.Println("exponenciacao =>", math.Pow(c, d))

}
