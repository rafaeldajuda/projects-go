# O Projeto: UrbanStay API

O projeto que vamos construir chama-se UrbanStay: uma API para uma plataforma de aluguel de imóveis por temporada.

Aqui está o nosso cronograma:

---

## Fase 1: O Alicerce (Domínio e CRUD Base)

### Requisitos da Fase 1:

1. Entidade Imóvel (Property): Deve conter ID, Nome, Descrição, Preço por noite e Endereço.
2. Regra de Negócio: Um imóvel não pode ser cadastrado com preço negativo ou nome vazio.
3. Casos de Uso (Usecases): 
    * Cadastrar um novo imóvel.
    * Listar todos os imóveis disponíveis.
4. Interface de Repositório: Defina como o domínio espera salvar esses dados (mesmo que a implementação inicial seja em memória).
5. Entrega (Delivery): Uma interface simples via HTTP (pode usar o pacote net/http padrão ou algo leve) para expor esses endpoints.

## Fase 2: Persistência e Configuração

### Requisitos da Fase 2:

1. Banco de Dados: Implemente um novo repositório em internal/repository/sqlite ou postgres (recomendo SQLite para facilitar seus testes locais agora).
2. Configuração: Crie um arquivo .env para armazenar a string de conexão do banco e a porta do servidor. Use uma biblioteca como godotenv ou viper.
3. Migrations (Opcional, mas recomendado): O sistema deve ser capaz de criar a tabela properties ao iniciar.
4. O Desafio SOLID: Você deve conseguir trocar o repositório de memória pelo de banco de dados no main.go alterando apenas uma linha de código, sem tocar no UseCase ou no Handler.

## Fase 3: Observabilidade e Robustez

### Requisitos da Fase 3:

1. Logs Estruturados: Substitua o log.Fatal por um logger profissional (como zerolog ou zap).
2. Configuração Dinâmica: Atualmente o caminho do banco está "hardcoded" no main. Vamos usar um arquivo .env para definir a string de conexão e a porta do servidor.
3. Graceful Shutdown: O seu main.go precisa saber como desligar o servidor e fechar a conexão com o banco de dados de forma limpa quando você der um Ctrl+C.