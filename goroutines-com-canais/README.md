# EXERCÍCIO 2: INTERMEDIÁRIO - GOROUTINES COM CANAIS

Objetivo: Usar goroutines com canais (channels) para sincronizar a comunicação entre elas.

**Descrição do Exercício:**

1. Crie uma função chamada calculateSquare que aceita dois parâmetros: um número int e um canal de comunicação (chan int) para enviar o resultado.
2. Dentro da função, calcule o quadrado do número e envie o resultado pelo canal.
3. No main, inicie um slice de números de 1 a 5.
4. Para cada número no slice, inicie uma goroutine que chama calculateSquare, passando o número e o canal.
5. Após iniciar todas as goroutines, receba e some os resultados dos quadrados de todas as goroutines a partir do canal.
6. Por fim, imprima o valor total da soma dos quadrados.
Exemplo de Saída Esperada:

A soma dos quadrados de 1 a 5 é 55.
```sh
A soma dos quadrados é: 55
```

**Dica:** Use make(chan int) para criar o canal e close(chan) para fechá-lo quando todos os valores forem recebidos.
