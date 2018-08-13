package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9

	// aredonda pra baixo
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Println("teste " + string(97))

	// int para string
	fmt.Println("teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)
	fmt.Println(reflect.TypeOf(num))

	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("verdadeiro")
	}
}
