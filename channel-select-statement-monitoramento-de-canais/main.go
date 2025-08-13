package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	// envia uma mensagem para o channel1 a cada 2 segundos
	go func() {
		for {
			time.Sleep(2 * time.Second)
			channel1 <- "Enviando para msg1"
		}
	}()

	// envia uma mensagem para o channel2 a cada 3 segundos
	go func() {
		for {
			time.Sleep(3 * time.Second)
			channel1 <- "Enviando para msg2"
		}
	}()

	count := 0 // contador para controlar o encerramento do channel
	for count < 5 {
		select {
		case msg1 := <-channel1: // recebendo mensagem do channel1
			fmt.Println("msg1 recebida:", msg1)
		case msg2 := <-channel2: // recebendo mensagem do channel2
			fmt.Println("msg2 recebida:", msg2)
		}
		count++
	}

	fmt.Println("fim do programa")
}
