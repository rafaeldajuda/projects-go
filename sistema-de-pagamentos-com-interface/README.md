# Exercício 2: Nível Intermediário – Sistema de Pagamentos com Interface

**Descrição:**

Crie uma interface chamada PaymentMethod com um método Pay(amount float64) string, que será implementada por diferentes métodos de pagamento, como CreditCard, Paypal e Bitcoin. Cada método de pagamento deve ter uma implementação diferente do método Pay, retornando uma string que descreva como o pagamento foi realizado. Em seguida, escreva uma função que processe todos os tipos de pagamento, aceitando uma lista de métodos de pagamento e tentando processar cada um até que todos sejam bem-sucedidos.

**Dicas:**
* A função Pay deve retornar uma mensagem dizendo qual método foi usado para o pagamento.
* Use uma lista de PaymentMethod e simule a tentativa de pagamento com cada um dos métodos.