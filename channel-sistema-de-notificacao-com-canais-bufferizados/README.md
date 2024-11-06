# EXERCÍCIO 2: SISTEMA DE NOTIFICAÇÃO COM CANAIS BUFFERIZADOS (NÍVEL INTERMEDIÁRIO)

**Objetivo:** Crie um programa que simule um sistema de notificações onde várias goroutines representam sensores que enviam notificações para o canal. A função principal (main) deve monitorar o canal de notificações e exibir as mensagens enviadas.

**Instruções:**
1. Crie um canal bufferizado para strings com capacidade de 3 mensagens.
2. Inicie três goroutines, cada uma representando um sensor:
    * O primeiro sensor deve enviar a mensagem "Sensor 1: Alerta de temperatura alta".
    * O segundo sensor deve enviar a mensagem "Sensor 2: Alerta de umidade baixa".
    * O terceiro sensor deve enviar a mensagem "Sensor 3: Alerta de pressão alta".
3. A função main deve monitorar o canal de notificações e exibir as mensagens enviadas por cada sensor até que todas as goroutines tenham enviado suas mensagens.
4. Use um sleep aleatório em cada goroutine para simular o envio de mensagens em tempos diferentes.
5. Feche o canal quando todas as goroutines concluírem suas operações.

**Dica:** O canal bufferizado permite que as goroutines enviem mensagens sem bloqueio imediato.

**Exemplo de Saída Esperada:**
```sh
Sensor 1: Alerta de temperatura alta
Sensor 2: Alerta de umidade baixa
Sensor 3: Alerta de pressão alta
```

**Bônus:** Adicione mais sensores e aumente o buffer do canal para ver como o sistema lida com múltiplas notificações.
