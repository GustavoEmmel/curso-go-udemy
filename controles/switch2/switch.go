package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { // switch true vai procurar o primeiro verdadeiro
	case t.Hour() < 12:
		fmt.Println("bom dia")
	case t.Hour() < 18:
		fmt.Println("boa tarde")
	default:
		fmt.Println("boa noite")
	}

	// desafio do switch
	fmt.Println(notaParaConceito(9.5))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))

}

func notaParaConceito(nota float64) string {

	/*
		var n = int(nota)

			switch {
			case n == 10:
				fallthrough
			case n >= 9:
				return "A"
			case n >= 8:
				return "B"
			case n >= 7:
				fallthrough
			case n >= 6:
				fallthrough
			case n >= 5:
				return "C"
			case n >= 4:
				fallthrough
			case n >= 3:
				return "D"
			default:
				return "E"
			}
	*/
	switch {
	case nota >= 9 && nota <= 10:
		return "A"
	case nota >= 8 && nota < 9:
		return "B"
	case nota >= 5 && nota < 8:
		return "C"
	case nota >= 3 && nota < 5:
		return "D"
	default:
		return "E"
	}

}
