# 🚀 Desafio: Simulador de Caixa Eletrônico (ATM)

Implemente um programa em Go que simule o funcionamento básico de um caixa eletrônico no terminal.

**Regras:**

1. O caixa deve ter um saldo inicial (por exemplo, R$ 1000).
2. O usuário deve poder executar os comandos:
    * saldo → mostra o saldo atual.
    * depositar <valor> → adiciona dinheiro ao saldo.
    * sacar <valor> → retira dinheiro do saldo (apenas se houver saldo suficiente).
    * extrato → mostra todas as operações realizadas até o momento.
    * sair → encerra o programa.

3. As operações devem ser registradas em uma lista (ex.: slice de strings) no formato:
    ```
    [Deposito: 200]
    [Saque: 100]
    [Saldo: 1100]
    ```
4. O programa deve rodar em loop interativo até o usuário digitar sair.

**Exemplo de execução:**
```
Comandos: saldo, depositar <valor>, sacar <valor>, extrato, sair
> saldo
Saldo atual: R$1000
> depositar 500
Depósito de R$500 realizado com sucesso
> sacar 200
Saque de R$200 realizado com sucesso
> saldo
Saldo atual: R$1300
> extrato
[Deposito: 500]
[Saque: 200]
> sair
Encerrando programa...
```

**🎯 Objetivos do desafio:**

* Praticar leitura de input (bufio.Reader ou fmt.Scanln).
* Trabalhar com condições (validar saldo suficiente).
* Usar slices para guardar o histórico de operações.
* Manter um loop principal de execução.