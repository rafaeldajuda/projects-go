# Exercício 2: Nível Intermediário – Função de Divisão com Tratamento de Erros Personalizado

**Descrição:**

Implemente uma função chamada Divide que receba dois números float64 e retorne o resultado da divisão. Caso o divisor seja zero, a função deve retornar um erro personalizado, indicando que não é possível dividir por zero. Crie um tipo de erro personalizado e utilize-o para retornar uma mensagem detalhada.

**Dicas:**

* Crie um tipo de erro específico para divisão por zero, implementando o método Error().
* A função Divide deve retornar dois valores: o resultado e um erro.
