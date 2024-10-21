**Goroutines**

Goroutines são funções ou métodos executados em concorrências. Podemos pensar nelas como uma espécie de **lightweight thread** que são gerenciadas pelo runtime do GO.

Chamamos de **lightweight thread** pois o custo para sua criação é muito menor quando comparada com um thread de verdade. Outro ponto positivo é que o runtime consegue aumentar ou diminuir a quantide de goroutines de acordo com a necessidade de aplicação, enquanto o número de thread normalmente é fixo.

## Características das Goroutines:

1. **Leves e eficientes:**
Goroutines são muito mais leves que threads tradicionais do sistema operacional. Enquanto threads podem consumir mais recursos, você pode iniciar centenas de milhares de goroutines sem sobrecarregar o sistema.

2. **Executadas de forma assíncrona:**
Uma goroutine pode ser iniciada e executada em paralelo com o restante do código. Não há garantia de quando ela vai ser agendada para rodar, mas o runtime do Go cuida de gerenciar essas goroutines de forma eficiente.

3. **Sintaxe simples:**
Para iniciar uma goroutine, basta prefixar a chamada de uma função com a palavra-chave go.

4. **Tudo é goroutines:**
Basicamento tudo em Go funciona com goroutines, inclusive a execução principal do programa. Por isso, qualquer chamada de uma nova goroutine acontece dentro de uma goroutine.

## Referência Utilizada

* O projeto foi desenvolvido com base no conteúdo disponível no artigo: [O que são e como funcionam as Goroutines](https://aprendagolang.com.br/o-que-sao-e-como-funcionam-as-goroutines/).
* Parte do texto também foi baseado e criado pelo ChatGPT.