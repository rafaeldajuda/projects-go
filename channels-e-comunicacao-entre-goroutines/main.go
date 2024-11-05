package main

import (
	"fmt"
	"time"
)

// Envia dados para o canal ch
func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

// Recebe dados do canal ch
func getData(ch chan string) {
	var input string
	for {
		input = <-ch // recebendo dados do canal ch
		fmt.Printf("%s\n", input)
	}
	close(ch) // fechando o canal
}

func main() {
	ch := make(chan string)
	go sendData(ch) // chamando a goroutine que envia dados
	go getData(ch)  // chamando a goroutine que recebe dados
	time.Sleep(1e9) // chamando a função Sleep para que de tempo da goroutines executem
}
