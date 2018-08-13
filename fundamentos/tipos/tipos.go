package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("literal inteiro eh", reflect.TypeOf(32000))

	// sem sinal (somente postivos) ... uint8 unit16 uint32 unit64

	var b byte = 255
	fmt.Println("o byte eh", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("o valor maximo do inte eh", i1)
	fmt.Println("o tipo do i1 eh", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela unicode (int32)
	fmt.Println("o rune eh", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("o tipo de x eh", reflect.TypeOf(x))
	fmt.Println("o tipo do literal 49.99", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("o tipo de bo eh", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Ola meu nome eh Gustavo"
	fmt.Println(s1 + "!")
	fmt.Println("o tamanho da string eh", len(s1))

	// string com multitplas linhas
	s2 := `Ola
meu
nome
eh
Gustavo`

	fmt.Println(s2 + "!")
	fmt.Println("o tamanho da string eh", len(s2))

	// char ???
	// var x char = 'b'
	char := 'a'
	fmt.Println("o tipo de char eh", reflect.TypeOf(char))
	fmt.Println(char)

}
