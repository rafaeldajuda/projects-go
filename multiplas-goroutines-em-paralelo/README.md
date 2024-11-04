# EXERCÍCIO 1: BÁSICO - MÚLTIPLAS GOROUTINES EM PARALELO

Objetivo: Iniciar e executar múltiplas goroutines em paralelo.

**Descrição do Exercício:**

1. Crie uma função chamada printNumbers que aceita um argumento id do tipo int.
2. Dentro da função, faça com que ela imprima o número do id seguido de uma contagem de 1 a 5. Por exemplo, se id for 1, a função deve imprimir "Goroutine 1: 1", "Goroutine 1: 2", etc.
3. No main, inicie 3 goroutines, cada uma chamando printNumbers com um valor diferente para id (1, 2 e 3).
4. Use time.Sleep no main para garantir que todas as goroutines tenham tempo de terminar.
Exemplo de Saída Esperada:

A saída pode variar em ordem, mas deve incluir as contagens de cada goroutine:

```sh
Goroutine 1: 1
Goroutine 1: 2
Goroutine 1: 3
Goroutine 1: 4
Goroutine 1: 5
Goroutine 2: 1
Goroutine 2: 2
Goroutine 2: 3
Goroutine 2: 4
Goroutine 2: 5
Goroutine 3: 1
Goroutine 3: 2
Goroutine 3: 3
Goroutine 3: 4
Goroutine 3: 5
```