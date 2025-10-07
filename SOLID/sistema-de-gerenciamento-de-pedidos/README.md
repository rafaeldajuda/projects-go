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
