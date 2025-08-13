# 🧠 Problema: Contador de Palavras Únicas

## Descrição:

Escreva um programa em Go que leia uma string (pode ser digitada diretamente no código ou lida do terminal) e conte quantas palavras únicas existem nela, desconsiderando letras maiúsculas/minúsculas e sinais de pontuação.

## Regras:

* Palavras devem ser consideradas iguais mesmo que estejam em letras diferentes (ex: "Go", "go", "GO" → mesma palavra).
* Palavras com pontuação no final devem ser tratadas corretamente (ex: "gopher," e "gopher" → mesma palavra).
* A saída deve ser o número total de palavras únicas.

**Exemplo:**

```sh
Input:
"Go is fun. Go, go, GO!"

Output:
3
```

Neste exemplo, as palavras únicas são: go, is, fun.