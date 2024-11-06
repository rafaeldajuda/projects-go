package main

import (
	"fmt"
	"time"
)

// envia os números de 1 a 5 para o canal 'ch'
func oneToFive(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
}

// envia os números de 6 a 10 para o canal 'ch'
func sixToTen(ch chan int) {
	ch <- 6
	ch <- 7
	ch <- 8
	ch <- 9
	ch <- 10
}

func main() {
	// cria o canal 'ch'
	ch := make(chan int)

	// chamando as goroutines que enviam números para o canal
	go oneToFive(ch)
	go sixToTen(ch)

	go func() {
		time.Sleep(time.Second) // 1 segundo de espera para dar tempo das goroutines executarem
		close(ch)               // fecha o canal
	}()

	// soma os números que o canal 'ch' recebe das goroutines
	soma := 0
	for range 10 {
		soma += <-ch
	}

	// imprime a soma dos números enviados pelas goroutines
	fmt.Println("A soma total é:", soma)
}
