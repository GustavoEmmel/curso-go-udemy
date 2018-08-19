package main

import (
	"fmt"
)

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("f2: %s %s\n", p1, p2)
}

func f3() string {
	return "f3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("f4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "ola", "mundo"
}

func main() {
	f1()
	f2("hello", "world")

	r3, r4 := f3(), f4("par 1", "par 2")
	fmt.Println(r3)
	fmt.Println(r4)

	// nao pode receber somente um, pode ignorar um deles atribuindo a um _
	r5_1, r5_2 := f5()
	fmt.Println(r5_1, r5_2)

	// assim estamos ignorando
	_, r5_2_1 := f5()
	fmt.Println(r5_2_1)
}
