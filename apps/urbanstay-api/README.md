# O Projeto: UrbanStay API

O projeto que vamos construir chama-se UrbanStay: uma API para uma plataforma de aluguel de imóveis por temporada.

Aqui está o nosso cronograma:

---

## Fase 1: O Alicerce (Domínio e CRUD Base)
Nesta fase, o foco é a pureza do domínio e a estruturação das camadas. Não quero me preocupar com bancos de dados complexos ou frameworks externos agora. Quero ver a lógica de negócio protegida.

### Requisitos da Fase 1:

1. Entidade Imóvel (Property): Deve conter ID, Nome, Descrição, Preço por noite e Endereço.
2. Regra de Negócio: Um imóvel não pode ser cadastrado com preço negativo ou nome vazio.
3. Casos de Uso (Usecases): 
    * Cadastrar um novo imóvel.
    * Listar todos os imóveis disponíveis.
4. Interface de Repositório: Defina como o domínio espera salvar esses dados (mesmo que a implementação inicial seja em memória).
5. Entrega (Delivery): Uma interface simples via HTTP (pode usar o pacote net/http padrão ou algo leve) para expor esses endpoints.

### O que espero ver na avaliação:

* Separação clara de pastas (Ex: internal/domain, internal/usecase, internal/repository, internal/delivery/http).
* Uso de Interfaces para garantir o desacoplamento.
* Injeção de dependência via construtores.
* Nenhuma lógica de banco de dados ou HTTP "vazando" para dentro do Usecase.

## Fase 2: Persistência e Configuração
Nesta fase, vamos tirar as "rodinhas de treinamento" da memória e usar um banco de dados real.

### Requisitos da Fase 2:

1. Banco de Dados: Implemente um novo repositório em internal/repository/sqlite ou postgres (recomendo SQLite para facilitar seus testes locais agora).
2. Configuração: Crie um arquivo .env para armazenar a string de conexão do banco e a porta do servidor. Use uma biblioteca como godotenv ou viper.
3. Migrations (Opcional, mas recomendado): O sistema deve ser capaz de criar a tabela properties ao iniciar.
4. O Desafio SOLID: Você deve conseguir trocar o repositório de memória pelo de banco de dados no main.go alterando apenas uma linha de código, sem tocar no UseCase ou no Handler.
