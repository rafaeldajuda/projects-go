# CHANNEL - SELECT STATEMENT

A cláusula select é uma estrutura que permite que uma goroutine aguarde e opere múltiplos channels de forma não bloqueante. Ela bloqueia a execução da função até que um dos channels esteja pronto para ser executado. Caso mais de um channel esteja pronto para ser executado, ela selecionará de forma aleatória qual executar. Funciona de forma semelhante a uma **switch**, mas é usado exclusivamente para operar canais.

## Como o **select** Funciona

Dentro de um **select**, cada "case" corresponde a uma operação de envio ou recebimento em um canal. Quando o **select** é executado, ele:

1. Avalia todas as operações de canais especificamente em seus "cases".
2. Escolhe aleátoriamente um dos "cases" prontos e o executa (ou seja, uma operação que não esteja bloqueada).
3. Se não houver canais prontos, o **select** bloqueia a execução até que um dos "cases" esteja pronto.

Isso torna o **select** extremamente útil para cenários onde várias operações de canais precisam ser monitoradas simultaneamente, permitindo que o programa responda ao primeiro canal pronto sem precisar especificar a ordem.

## Sintaxe Básica

```go
select {
case msg1 := <- channel1:
    // Processa msg1 recebida de channel1
case channel2 <- msg2:
    // Envia msg2 para channel2
default:
    // Executa se se nenhum canal estiver pronto (opcional)
}
```

* Cada **case** pode ser uma operação de recebimento **(<-channel1)** ou de envio **(channel2 <- msg2)**.
* O bloco **default** é executado se nenhum canal estiver pronto, evitando o bloqueio do **select**.

## Exemplo

```go
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
			fmt.Println(<-total)    // Recebe e imprime o valor do canal 'total'
		}
		exit <- true    // Envia sinal de saída após consumir 10 valores
	}()

    // Chama a função 'sum', que inicia o envio de valores ao canal 'total'
	sum(total, exit)
}
```

* A função **sum()** vai ficar somando números aleátórios (de 0 a 19) infinitamente. Essa função receberá 2 channels por parâmetros, sendo um para comunicar o valor atual da somatória e outro para receber o sinal de exit.

    ```go
    func sum(total chan int, exit chan bool) {
        valor := rand.Intn(20)
        for {
            select {
            case total <- valor:
                valor += rand.Intn(20)
            case <-exit:
                fmt.Println("exit :)")
                return
            }
        }
    }
    ```

    Mesmo a função tendo um loop infinito, o channel **exit** irá controlar o fim da sua execução, já que quando ele receber um valor, o **return** será executado, matando a execução do loop.

* Na função principal, criamos uma goroutine anônima para controlar a execução da função **sum()**. Como limitamos o loop em 10 iterações e a cláusula **select** bloqueia a execução até que um dos channels esteja pronto, o número de voltas do loop da função **sum()** deve ser parecido com o da nossa função anônima.

    ```go
    func main() {
        total := make(chan int)
        exit := make(chan bool)

        go func() {
            for i := 0; i < 10; i++ {
                fmt.Println(<-total) 
            }
            exit <- true 
        }()

        sum(total, exit)
    }
    ```

    **OBS:** Outro ponto de atenção na função principal, é que o channel **total** não tem buffer. Dessa forma, a soma do próximo número aleatório só irá acontecer após alguma outra rotina ler a última soma executada. 

## Quando Usar o **select**

* **Concorrência entre canais:** Para gerenciar operações concorrentes de multíplos canais e escolher o primeiro canal pronto.
* **Time-out:** Usar **select** com um timer para opereções limitadas de tempo, úitil para evitar bloqueios indefinidos.
* **Não-bloquei com default:** Para realizar opereções não-bloqueantes, usando o **default** com fallback se nenhum canal estiver pronto.

## Exemplo com Time-out

Exemplo com um time-out usando o pacote **time.After**, que cria um canal que recebe um valor após um tempo determinado:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    messages := make(chan string)

    // Goroutine que envia uma mensagem após 2 segundos
    go func() {
        time.Sleep(2 * time.Second)
        messages <- "Mensagem recebida!"
    }()

    select {
    case msg := <-messages:
        fmt.Println(msg)
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout: Nenhuma mensagem recebida em 1 segundo")
    }
}
```

Saída:
```sh
Timeout: Nenhuma mensagem recebida em 1 segundo
```

No exemplo acima, o **select** espera uma mensagem do canal **messages** ou de **time.After**. Como o tempo limite de **time.After** ocorre antes que a mensagem seja enviada, ele exibe "Timeout: Nenhuma mensagem recebida em 1 segundo".

## Referências Utilizadas
* O projeto foi desenvolvido com base no conteúdo disponível no artigo: [Múltiplos channels e a cláusula select](https://aprendagolang.com.br/multiplos-channels-e-a-clausula-select/#:~:text=A%20cl%C3%A1usula%20select%20%C3%A9%20utilizada,de%20forma%20aleat%C3%B3ria%20qual%20executar.)
* Parte do texto também foi baseado e criado pelo ChatGPT
