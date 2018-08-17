package main

import (
	"fmt"
	"time"
)

// recebendo parametro generico
func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteito"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "funcao"
	default:
		return "nao sei"
	}
}

func main() {
	fmt.Println(tipo(2.4))
	fmt.Println(tipo(1))
	fmt.Println(tipo("gustavo"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now))
}
