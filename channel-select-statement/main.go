package main

import (
	"fmt"
	"math/rand"
)

// Comentários feito pelo ChatGPT

// Função que envia valores incrementados aleatoriamente para o canal 'total'.
// Também verifica o canal 'exit' para saber se deve encerrar sua execução.
func sum(total chan int, exit chan bool) {
	// Inicializa um valor aleatório entre 0 e 19
	valor := rand.Intn(20)
	for {
		select {
		// Caso o canal 'total' esteja pronto para receber, envia o valor atual
		case total <- valor:
			// Incrementa 'valor' com um novo número aleatório entre 0 e 19
			valor += rand.Intn(20)
			// Caso receba uma sinalização de saída pelo canal 'exit', imprime uma mensagem e encerra
		case <-exit:
			fmt.Println("exit :)")
			return
		}
	}
}

func main() {
	// Canal para envio de valores inteiros
	total := make(chan int)
	// Canal para sinalização de saída
	exit := make(chan bool)

	// Goroutine que consome 10 valores do canal 'total' e sinaliza saída ao finalizar
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-total) // Recebe e imprime o valor do canal 'total'
		}
		exit <- true // Envia sinal de saída após consumir 10 valores
	}()

	// Chama a função 'sum', que inicia o envio de valores ao canal 'total'
	sum(total, exit)
}
