package main

import (
	"fmt"
)

func main() {
	i := 1

	// Para criar um ponteiro usar a anotacao *
	var p *int = nil

	p = &i //pegando o endereco de memoria i e atribuindo ao ponteiro

	// Go nao tem aritimetica de ponteiros
	*p++ //desaciou o valor do ponteiro (pegou o valor) e incrementou o valor da variavel
	i++

	fmt.Println(p, *p, i, &i)
}
