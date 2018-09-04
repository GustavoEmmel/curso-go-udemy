package matematica

import (
	"testing"
)

const erroPadrao = "valor esperado %v, mas o resultado econtrado foi %v"

func TestMedia(t *testing.T) {
	t.Parallel()
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

// os metodos de testes devem respeitar esse padrao.
// arquivo terminando com _test.go
// metodo de teste initiando com TestMethod e esperando sempre um t *testing.T
// iniciando todos os testes com t.Parallel() fara todos testes rodarem de maneira paralela
// para rodar todos os testes, executar o comando abaixo
// C:\Users\Windows10\go\src\curso-go-udemy\testes> go test ./...
// para pegar a cobertura de testes:
// C:\Users\Windows10\go\src\github.com\matematica> go test -cover
// PASS
// coverage: 100.0% of statements
// ok      github.com/matematica   0.061s
