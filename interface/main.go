package main

import (
	"fmt"
)

// Interface Animal que define o contrato para qualquer tipo que implemente a função Onomatopeia.
// Qualquer struct que implementar essa função será considerado do tipo Animal.
type Animal interface {
	Onomatopeia() string
}

// struct Urso que implementa a função Onomatopeia de forma específica para a espécie.
type Urso struct{}

// Onomatopeia retorna o som característico do urso.
func (u Urso) Onomatopeia() string {
	return "Roarrrrr"
}

// struct Tucano que implementa a função Onomatopeia de forma específica para a espécie.
type Tucano struct{}

// Onomatopeia retorna o som característico do tucano.
func (t Tucano) Onomatopeia() string {
	return "Tri tri tri"
}

// Função Play recebe um parâmetro do tipo Animal (qualquer struct que implemente a interface Animal).
// A função chama o método Onomatopeia para imprimir o som do animal.
func Play(animal Animal) {
	fmt.Println(animal.Onomatopeia())
}

func main() {
	// Cria uma instância de Urso e passa para a função Play, que imprime o som do urso.
	urso := Urso{}
	Play(urso)

	// Cria uma instância de Tucano e passa para a função Play, que imprime o som do tucano.
	tucano := Tucano{}
	Play(tucano)
}
