package main

import (
	"math"
)

// inicializando com letra maiuscula eh PUBLICO
// (visibilidade dentro e fora do pacote)
// as structs e os metodos devem ter comentarios na linha acima

// inicializando com letra minuscula eh PRIVADO
// (visibilidade apenas dentro do pacote)

// por exemplo...
// Ponto - gerara algo publico
// ponto ou _Ponto - gerara algo privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

// estamos usando o retorno limpo, entao nao precisa dizer o que retornar
func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)

	return
}

// Distancia eh responsavel por calcular a distancia linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
