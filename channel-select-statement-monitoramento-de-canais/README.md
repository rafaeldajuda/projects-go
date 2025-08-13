# Exercício 1: Monitorando dois canais (Nível Iniciante)

Crie um programa que tenha duas goroutines enviando mensagens para dois canais diferentes (channel1 e channel2). A primeira goroutine deve enviar uma mensagem para channel1 a cada 2 segundos, enquanto a segunda deve enviar uma mensagem para channel2 a cada 3 segundos.

Use um select statement na goroutine principal para monitorar os dois canais e imprimir as mensagens conforme elas chegam. Finalize o programa após receber 5 mensagens no total (de ambos os canais).

**Dicas:**

* Utilize um contador para acompanhar quantas mensagens já foram recebidas.
* O select deve lidar com mensagens de ambos os canais.
