# CHANNELS E COMUNICAÇÃO ENTRE GOROUTINES

Channels, ou canais, são estruturas de comunicação que permitem trocas de informações entre goroutines. Utilizar channels é a melhor maneira de se utilizar as goroutines, permitindo trocas de informações de maneira mais segura e sincronizada, não precisando fazer travas explicitadas que podem ser muito complexas de se fazer e propensas a erros.

## Iniciando um canal

```go
var nome chan tipo
```

* Ao declarar um canal, o valor inicial é **nil**
* Um canal só pode transmitir dados de um tipo (chan string, chan int, ...)
* Um canal é uma fila de FIFO (first in first out)
* Um canal é uma referência, o que se significa que precisamos usar o **make** para alocar memória para ele.
* Canais são objetos de primeira classe, eles podem ser armazenados em variáveis, passados como argumentos, retornados de funções e enviados por canais.

**Como declarar um canal**

```go
var ch1 chan string
ch1 = make(chan string)
```

```go
ch1 := make(chan string)
```

## Fechando um canal

É uma boa prática fechar canais após utiliza-los.

```go
close(ch1)
```

## Operador de comunicação

Os canais possui um operador de comunição de transmissão de dados entre eles. É um operador bem intuitivo que indica o fluxo da comunição (receber/enviar).

**Envio de dados**

* A sintaxe para enviar um valor para o canal é **ch <- valor**.
* O canal ficará bloqueado até que outra goroutine esteja pronta para receber o valor.

```go
ch <- 42
```

**Recebimento de dados**

* A sintaxe para receber um valor de uma canal é **valor := <- ch**, ou caso a variável já esteja declarada, **valor = <- ch**.

```go
num := <- ch 
```

* Uma outra forma de receber o valor é **<- ch**, pode ser usada para obter próximo valor do canal. O valor será descartado, porém pode ser testado.

```go
if <- ch != 1000 {
    ...
}
```

---

O mesmo operador **<-** é usado para enviar e receber, mas o Go descrobre o que fazer dependendo dos operandos.

**Comunicação e sincronização**

Os canais são bloqueantes por padrão. Isso significa que:

* Quando uma goroutine envia um valor a um canal, ela fica bloqueada até que outra goroutine leia o valor.
* Quando uma goroutine tenta ler de um canal vazio, ela fica bloqueada até que um valor esteja disponível.

Esse comportamento garante uma sincronização natural, o que é especialmente útilpara coordenar o fluxo de execução de goroutines.

## Exemplo de comunicação entre goroutines

```go
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
	time.Sleep(1e9) // chamando a função Sleep para que de tempo das goroutines executem
}
```

**Observações importantes:**

* A função **main()** aguarda 1 segundo para permitir que ambas as goroutines sejam concluídas.
* **getData()** trabalha com um loop infinito, que é interrompido somente quando **sendData** termina e o canal **ch** está vazio.
* Caso sejam removidas uma ou todas as palavras-chaves go, o programa deixará de funcionar e o tempo do 
Go lançará um erro: 
    **Error run <path> with code Crashed Fatal error: all goroutines are asleep - deadlock!**

Essa ocorrência se dá polo fato de o tempo de execução do Go, ser capaz de detectar que todas as goroutines estão aguardando por algo, impedindo assim o avanço do programa. Trata-se de uma forma de impasse (deadlock) que o tempo de execução é capaz de identificar.

## Referência Utilizada
* O projeto foi desenvolvido com base no conteúdo disponível no artigo: [Concorrência em Go](https://medium.com/@rafaelmgr12/concorr%C3%AAncia-em-go-85b6a127b12f#:~:text=Os%20channels%20s%C3%A3o%20canais%20de,corrida%20e%20problemas%20de%20concorr%C3%AAncia.).
* Parte do texto também foi baseado e criado pelo ChatGPT.