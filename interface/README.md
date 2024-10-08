# Projeto: Entendendo Interfaces em Go

Este projeto tem como objetivo estudar e compreender o funcionamento de **interfaces** na linguagem Go (Golang). 

### Propósito

O código exemplifica como utilizar interfaces em Go para permitir que diferentes tipos (structs) possam compartilhar o mesmo comportamento através de uma função comum. No caso deste exemplo, a interface `Animal` define um método chamado `Onomatopeia`, e dois structs (`Urso` e `Tucano`) implementam esse método de forma específica, retornando sons característicos de cada animal. 

A função `Play` recebe qualquer tipo que implemente a interface `Animal` e imprime o som característico do animal, demonstrando a flexibilidade e o poder das interfaces em Go.

### Estrutura do Projeto

- **Animal Interface**: Define um contrato para qualquer struct que implemente a função `Onomatopeia() string`.
- **Urso e Tucano Structs**: Implementam a função `Onomatopeia` retornando o som específico de cada animal.
- **Função Play**: Recebe como argumento um tipo que implementa a interface `Animal` e imprime o som do respectivo animal.

### Referência Utilizada

O projeto foi desenvolvido com base no conteúdo disponível no artigo: [Trabalhando com Interfaces](https://aprendagolang.com.br/trabalhando-com-interfaces/).

**OBS:** README feito no ChatGTP.
