# 🚀 Desafio: Gerenciador de Tarefas no Terminal

Implemente um pequeno programa em Go que funcione como um gerenciador de tarefas no terminal.

**Regras:**

* Cada tarefa deve ter:
    * Um ID único (inteiro).
    * Um Título (string).
    * Um campo Concluída (bool).

* O programa deve permitir:
    * Adicionar uma nova tarefa.
    * Listar todas as tarefas.
    * Marcar uma tarefa como concluída a partir do ID.
    * Remover uma tarefa pelo ID.

* Use um slice ou map para armazenar as tarefas em memória (não precisa salvar em arquivo).

* O programa deve rodar em loop até o usuário digitar "sair".

**Exemplo de execução:**
```txt
Comandos disponíveis: add, list, done, delete, sair
> add Estudar Go
Tarefa adicionada com ID 1
> add Fazer exercícios
Tarefa adicionada com ID 2
> list
1 - Estudar Go [pendente]
2 - Fazer exercícios [pendente]
> done 1
Tarefa 1 marcada como concluída
> list
1 - Estudar Go [concluída]
2 - Fazer exercícios [pendente]
> delete 2
Tarefa 2 removida
> list
1 - Estudar Go [concluída]
> sair
Encerrando programa...
```