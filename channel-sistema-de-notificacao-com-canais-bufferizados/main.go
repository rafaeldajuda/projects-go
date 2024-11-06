package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sensor(id int, msg string, notificaion chan string, done chan bool) {
	n := rand.Intn(7) + 1                                // Gera um número aleatório de 1 a 7
	time.Sleep(time.Second * time.Duration(n))           // Tempo de espera aleatório
	notificaion <- fmt.Sprintf("Sensor %d: %s", id, msg) // Envia a mensagem formatada para o canal de notificação
	done <- true                                         // Sinaliza que o sensor terminou sua execução
}

func main() {
	// Cria os canais 'notification' e 'done' com buffer para mensagens
	notification := make(chan string, 5)
	done := make(chan bool, 5)

	// Inicia cinco goroutines, cada uma representando um sensor que envia uma mensagem específica
	go sensor(1, "Alerta de temperatura alta", notification, done)
	go sensor(2, "Alerta de umidade baixa", notification, done)
	go sensor(3, "Alerta de pressão alta", notification, done)
	go sensor(4, "Alerta de temperatura baixa", notification, done)
	go sensor(5, "Alerta de umidade alta", notification, done)

	// Goroutine para fechar o canal de notificação quando todos os sensores terminarem
	go func() {
		for range 5 { // Aguarda cinco confirmações de conclusão no canal 'done'
			<-done
		}
		close(notification) // Fecha o canal 'notification'
	}()

	// // Lê as mensagens do canal de notificação e as imprime
	for msg := range notification {
		fmt.Println(msg)
	}

	fmt.Println("Todas notificações recebidas!")
}
