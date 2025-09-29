# Desafio GoLang - Gerenciador de Tarefas com API e MongoDB

## 🎯 Objetivo
Construir uma **API REST em Go** para gerenciar tarefas (to-do list), persistindo os dados em **MongoDB**.  

O desafio é projetado para nível **iniciante a intermediário** e deve ser concluído em **2 a 4 horas**.

---

## 📌 Requisitos

### Funcionalidades obrigatórias
1. **Criar uma tarefa**
   - Endpoint: `POST /tasks`
   - Body JSON exemplo:
     ```json
     {
       "title": "Estudar Go",
       "description": "Praticar structs, interfaces e goroutines",
       "completed": false
     }
     ```

2. **Listar todas as tarefas**
   - Endpoint: `GET /tasks`

3. **Buscar uma tarefa por ID**
   - Endpoint: `GET /tasks/{id}`

4. **Atualizar uma tarefa**
   - Endpoint: `PUT /tasks/{id}`  
   - Permitir atualizar `title`, `description` e `completed`.

5. **Remover uma tarefa**
   - Endpoint: `DELETE /tasks/{id}`

---

## 🔧 Tecnologias sugeridas
- **Go** (padrão `net/http` ou algum framework como `gin` ou `echo`)
- **MongoDB** (utilizar driver oficial `go.mongodb.org/mongo-driver`)
- **Docker** (opcional, para subir o MongoDB facilmente)

---

## 🚀 Extra (opcional, para ir além)
- Adicionar paginação no endpoint `GET /tasks`
- Implementar busca por título (`GET /tasks?title=...`)
- Criar testes unitários para os handlers principais
- Criar um `docker-compose.yml` que sobe a API e o MongoDB juntos

---

## 📝 O que será avaliado
- Organização do código (pacotes, nomes, clareza)
- Boas práticas com Go (uso de structs, handlers, context, etc.)
- Interação correta com o MongoDB
- Tratamento de erros e respostas adequadas da API
- Documentação (README atualizado)

---

## ▶️ Como rodar
1. Clone o repositório
2. Instale as dependências
   ```bash
   go mod tidy
 