# EXERCÍCIO 1: SOMADOR SIMPLES (NÍVEL INICIANTE)

**Objetivo:** Crie um programa onde duas goroutines enviem valores inteiros a um canal compartilhado. A função principal (main) deve ler esses valores e exibir a soma total.

**Instruções:**
1. Crie um canal do tipo int.
2. Inicie duas goroutines:
    * A primeira goroutine deve enviar os números 1 a 5 ao canal.
    * A segunda goroutine deve enviar os números 6 a 10 ao canal.
3. Na função main, leia os valores do canal usando um loop e some esses valores até que todas as goroutines tenham terminado de enviar.
4. Exiba a soma total dos números enviados ao canal.

**Dica:** Feche o canal após as goroutines concluírem suas operações para que o loop de leitura finalize automaticamente.

**Exemplo de Saída Esperada:**
```sh
A soma total é: 55
```