package main

import (
	"fmt"
	"time"
)

// Função printNumbers que imprime os núemros de 1 a 5 com um ID
func printNumbers(id int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
	}
}

func main() {
	// Três goroutines que chama a função printNumbers com os IDs 1, 2 e 3
	go printNumbers(1)
	go printNumbers(2)
	go printNumbers(3)

	// Chamando a função Sleep para dar tempo das goroutes serem executadas
	time.Sleep(3 * time.Second)
}
