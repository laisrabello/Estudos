// METÓDOS EM STRUCTS

package main

import "fmt"

type Pessoa struct {
	Nome      string
	Idade     int
	Profissao string
}

func (p Pessoa) ListaNomeEIdade() string {
	return fmt.Sprintf("%s tem %d anos", p.Nome, p.Idade)
}

func main() {

	pessoa := Pessoa{
		Nome:      "Laís",
		Idade:     27,
		Profissao: "Analista",
	}

	println(pessoa.ListaNomeEIdade())
}
