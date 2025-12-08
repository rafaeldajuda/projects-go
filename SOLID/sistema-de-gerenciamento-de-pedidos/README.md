# 🧩 Desafio Golang — Order Management System (SOLID + MongoDB/Redis)

## 📘 Descrição do Desafio

Desenvolver uma aplicação em **Go (Golang)** que implemente um **sistema de gerenciamento de pedidos (Order Management System)** aplicando de forma prática os **princípios SOLID**.

O projeto deve utilizar **MongoDB ou Redis** como banco de dados e ser estruturado de forma **modular, extensível e orientada a interfaces**.

---

## 🎯 Objetivos

- Aplicar **todos os princípios SOLID** em um contexto real.
- Criar um sistema escalável e de fácil manutenção.
- Demonstrar **boas práticas de arquitetura** em Go.
- Utilizar **MongoDB** ou **Redis** como camada de persistência.

---

## 🧱 Funcionalidades obrigatórias

### 1. Criar Pedido
Cria um novo pedido com os dados do cliente e produtos.

### 2. Listar Pedidos
Retorna todos os pedidos armazenados no banco de dados.

### 3. Buscar Pedido
Busca um pedido pelo seu ID.

### 4. Atualizar Status
Atualiza o status de um pedido existente.

---

## 🧩 Estrutura esperada do Pedido

```json
{
  "id": "uuid-gerado",
  "customer": {
    "name": "João Silva",
    "email": "joao@email.com"
  },
  "products": [
    {
      "name": "Teclado Mecânico",
      "quantity": 1,
      "unit_price": 350.0
    },
    {
      "name": "Mouse Gamer",
      "quantity": 1,
      "unit_price": 200.0
    }
  ],
  "total_value": 550.0,
  "status": "CRIADO"
}
```

---

# Aplicação do SOLID no Projeto

**SOLID** reúne cinco princípios que tornam o desenvolvimento mais simples, facilitando a manutenção e evolução do software.
Eles são fundamentais na programação orientada a objetos e podem ser aplicados em qualquer linguagem que siga esse paradigma.

## S - Single Responsability Principle (SRP)

Cada classe deve ter somente uma responsabilidade com seus próprios métodos e atributos. Se a classe tiver um ou mais métodos/atributos não relacionado a função da classe, o ideal seria criar uma nova classe somente para esses métodos.

**Vantagens:**
- Facilidade para fazer manutenções 
- Reusabiliade das classes
- Facilidade para realizar testes
- Simplicidade da legibilidade do código

**No projeto:**
```text
├── cmd
└── internal
    ├── api           -> Lidar com o HTTP 
    ├── domain        -> Representar entidades e regras puras
    ├── repository    -> Persistência de dados (banco de dados)
    └── service       -> Regras de negócio e orquestração
```

Cada arquivo/pasta faz somente aquilo que pertence ao seu papel.

## O - Open/Closed Principle (OCP)

Componentes devem estar abertos para extenção, mas fechados para modificação. Isso significa que devemos conseguir adicionar novos comportamentos, mas sem alterar o código existe, diminuindo o risco de quebrar funcionalidades já estáveis.

O uso de abstrações, como interfaces e classes, ajuda a isolar comportamentos e adicionar novas funcionalidades sem aumentar a complexidade da classe original.

**No projeto:**
```text
├── repository
|   ├── orderRepository.go      -> interface
|   ├── mongoRepository.go      -> implementação
|   └── redisRepository.go      -> implementação
└── service                     -> recebe a interface
```

* Define uma interface.
* Implementa Mongo e Redis como estratégias diferentes.
* O service depende só da interface, não das implementações.
* Trocar Mongo por Redis, ou vice e versa, não irá alterar o service.
* Quer adicionar **postgreSQL**? Só precisa adicionar um novo arquivo **postgreRepository.go**.
* Você extende o sistema sem mexer no restante.

## L - Liskov Substituion Principle (LSP)

Uma classe-filha deve ser capaz de executar tudo que a sua classe-mãe faz. Esse princípio se conecta como polimorfismo e reforça esse pilar da POO.

**No projeto:**

orderRepository.go:
```go
type OrderRepository interface {
	Create(order domain.Order) (*domain.Order, error)
	ListAll() ([]domain.Order, error)
	FindByID(id string) (*domain.Order, error)
	Update(order domain.Order) (*domain.Order, error)
}
```

Tanto o **mongoRepository.go** e **redisRepository.go** implementam exatamente esse métodos. O resultado disso é que o **service.go** pode receber qualquer um.

service.go:
```go
func NewOrderService(repo repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}
```

* Mongo substitui Redis.
* Redis Substitui Mongo.
* Nenhum comportamento inesperado aparece.

Objetos podem ser substituidos por objetos por seus subtipos sem que isso afete a execução correta do programa.

## I - Interface Segregation Principle (ISP)

Interfaces devem ser pequenas e específicas, dependendo somente dos seus campos e métodos. Caso a interface fique grande ou começou a ter métodos que a princípio que não faz sentido como o propósito do objeto, então é melhor quebrar a criar duas ou mais interfaces. Devemos criar interfaces específicas ao invés de termos uma única interface generérica

**"Uma classe não deve ser forçada a implementar interfaces e métodos que não serão utilizados".**

* Promoce coesão e flexibilidade.
* Fácil manutenção.

## D - Dependency Inversion Principle (DIP)

Módulos de alto nível não devem depender de módulos de baixo nível, e sim, ambos devem depender de abstrações.
Abstrações também não devem depender de detalhes, mas detalhes devem depender de abstrações.

**No projeto:**

**Service** depende da interface, não da implementação.
```go
type OrderService struct {
	repo repository.OrderRepository
}
```

* O service **não sabe** se usa Mongo, Redis, PostgreSQL ou mock.
* Ele depende de uma **abstração** (OrderRepository).

**Quem escolhe a implementação?**

Nesse projeto: cmd/main.go
```go
// 1 - Conectar ao MongoDB
.
.
.
// 2 - Criar repositories
orderRepository := repository.NewMongoOrderRepository(collection.Database())

// 3 - Criar Services
orderService := service.NewOrderService(orderRepository)

// 4 - Criar API/hanlder
handler := api.NewHandler(orderService)
handler.StartHandler(fiber.Config{})
handler.Routes()
.
.
.
```

* Camadas altas (service) não dependem da infra.
* Ambos dependem de interfaces.
* Baixo nível (mongoRepository) implementa a interface.
