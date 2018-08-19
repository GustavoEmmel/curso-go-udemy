package main

import (
	"fmt"
	"reflect"
)

func main() {
	// array
	a1 := [3]int{1, 2, 3}
	//slice
	s1 := []int{1, 2, 3}

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// slice nao eh um array! define um pedaco de um array

	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:2] // novo slice apontando pro mesmo array
	fmt.Println(a2, s3)

	// voce pode imaginar um slice como: tamanho e ponteiro para um elemento de um array

	s4 := s2[:1]
	fmt.Println(a2, s2, s4)
}
