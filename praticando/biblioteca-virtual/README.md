# 🚀 Desafio: Biblioteca Virtual

Implemente um sistema de gerenciamento de livros no terminal.

**📌 Regras**

1. Cada livro deve ter:
    * Um ID único (string ou int).
    * Um Título (string).
    * Um Autor (string).
    * Um campo Disponível (bool).
2. O sistema deve oferecer os comandos:
    * add <titulo>;<autor> → adiciona um novo livro.
    * list → lista todos os livros (com status disponível/emprestado).
    * borrow <id> → marca o livro como emprestado (se estiver disponível).
    * return <id> → devolve o livro (marca como disponível novamente).
    * remove <id> → remove o livro do sistema.
    * exit → encerra o programa.
3. As operações devem ser armazenadas em memória (não precisa salvar em arquivo).
4. O programa deve rodar em loop interativo, aceitando comandos até exit.

**💻 Exemplo de execução**
```
Comandos: add <titulo>;<autor>, list, borrow <id>, return <id>, remove <id>, exit
> add Dom Casmurro;Machado de Assis
Livro adicionado com ID 1
> add A Hora da Estrela;Clarice Lispector
Livro adicionado com ID 2
> list
1 - Dom Casmurro (Machado de Assis) [Disponível]
2 - A Hora da Estrela (Clarice Lispector) [Disponível]
> borrow 1
Livro 1 emprestado com sucesso
> list
1 - Dom Casmurro (Machado de Assis) [Emprestado]
2 - A Hora da Estrela (Clarice Lispector) [Disponível]
> return 1
Livro 1 devolvido com sucesso
> remove 2
Livro 2 removido
> exit
Encerrando programa...
```

**🎯 Objetivos do desafio**
* Praticar structs (Livro).
* Trabalhar com slices ou map para armazenamento.
* Implementar validações (não emprestar livro já emprestado, não devolver livro disponível, etc.).
* Exercitar parsing de strings (add <titulo>;<autor>).
* Criar um loop principal e funções auxiliares para cada operação.