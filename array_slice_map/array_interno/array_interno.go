package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	s1[0] = 7
	fmt.Println(s1, s2)

	// criando um slice dinamico
	s3 := make([]int, 0, 0)
	s3 = append(s3, 1, 2, 3)
	fmt.Println(s3)
}
