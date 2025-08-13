# Exercício 2: Implementando Timeout com Select (Nível Intermediário)

Implemente um programa que simula a comunicação com dois servidores fictícios, representados por goroutines que enviam uma resposta aleatória para um canal após um atraso de 1 a 5 segundos. Utilize o pacote time.After para definir um timeout de 3 segundos para a resposta de cada servidor.

Se um servidor enviar uma resposta dentro do limite de tempo, exiba a mensagem recebida.
Caso contrário, exiba uma mensagem indicando que o servidor demorou demais para responder.
O programa deve encerrar após processar as respostas de ambos os servidores, com ou sem timeout.

**Dicas:**

* Use rand.Intn para simular o atraso aleatório dos servidores.
* Use select para aguardar as respostas dos servidores ou o timeout.
