package main

import (
	"fmt"
)

func main() {
	array1 := [3]int{1, 2, 3}

	// criando um slice dinamicamente
	var slice1 []int
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)

	// tem que passar um slice, nao pode passar um array
	copy(slice2, slice1)
	fmt.Println(slice2)
}
