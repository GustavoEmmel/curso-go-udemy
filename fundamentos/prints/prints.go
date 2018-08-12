package main

import (
	"fmt"
)

func main() {
	fmt.Print("mesma")
	fmt.Print(" linha")

	fmt.Println(" nova")
	fmt.Println("linha")

	x := 3.141516
	//fmt.Println("o valor de x eh " + x)
	xs := fmt.Sprint(x)
	fmt.Println("o valor de x eh " + xs)
	fmt.Println("o valor de x eh", x, "!!!")

	// estamos imprimindo um valor float com 2 casas decimais
	fmt.Printf("o valor de x e %.2f", x)
	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t, %s", a, b, b, c, d)

	// %v eh um generico para printar a variavel
	fmt.Printf("\n%v, %v %v %v", a, b, c, d)

}
