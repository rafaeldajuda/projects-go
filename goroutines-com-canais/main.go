package main

import (
	"fmt"
	"time"
)

// Função que calcula o quadrado de um número e o envia para o canal 'c'
func calculateSquare(num int, c chan int) {
	fmt.Println("Enviando:", num)
	c <- (num * num)
}

func main() {
	// Array de números para calcular o quadrado
	var numbers = [5]int{1, 2, 3, 4, 5}
	// Cria um canal para comunicação entre goroutine
	var c = make(chan int)

	// Inicia uma goroutine para calcular o quadrado de cada número no array
	for _, v := range numbers {
		go calculateSquare(v, c)
	}

	////////////////////////////////////////////////////////////////////////////////////////////

	// Fechar o canal de forma explícita após um tempo para que não cause deadlock ao tentar ler,
	// pois não se sabe quantos valores serão lidos.
	// Normalmente utilizado para ler todos os valores de um canal.
	go func() {
		time.Sleep(time.Second) // Espera 1 segundo para garantir que todos os valores sejam enviados
		close(c)                // Fecha o canal para indicar o fim dos envios
	}()

	// Lê valores do canal até que ele seja fechado, somando os valores recebidos
	soma := 0
	for num := range c {
		soma += num
	}

	////////////////////////////////////////////////////////////////////////////////////////////

	// Alternativa: Leitura de valores autómatica. Útil quando se sabe quantos valores serão lidos.
	// soma := 0
	// for range numbers {
	// 	soma += <-c // Recebe o valor do canal e soma
	// }

	////////////////////////////////////////////////////////////////////////////////////////////

	// Exibe o resultado final
	fmt.Println("A soma dos quadrados é:", soma)
}
