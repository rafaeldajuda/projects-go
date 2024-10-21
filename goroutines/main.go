package main

import (
	"fmt"
	"time"
)

// Função que imprime números de 0 a 9 com uma pausa de 200 milissegundos entre eles
func numeros() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 200) // Pausa de 200ms entre cada número
	}
}

// Função que imprime letras de 'a' até 'i' com uma pausa de 300 milissegundos entre elas
func letras() {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond * 300) // Pausa de 300ms entre cada letra
	}
}

func main() {
	// Executa a função numeros em uma nova goroutine (executa de forma concorrente)
	go numeros()
	// Executa a função letras em uma nova goroutine (executa de forma concorrente)
	go letras()

	// Pausa o programa principal por 5 segundos para permitir que as goroutines terminem
	time.Sleep(5 * time.Second)
	fmt.Printf("fim da execução\n") // Indica o fim da execução do programa principal
}
