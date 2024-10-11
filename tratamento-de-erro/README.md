# Tratamento de Erro

Em Go, **error** é uma interface padrão definida da seguinte forma:

```go
type error interface {
    Error() string
}
```

O tratamento de erros no Golang é um dos aspectos mais fundamentais e simples da linguagem. Golang usa um mecanismo explícito para lidar com erros, diferentemente de muitas outras linguagens que usam exceções (exceptions). Isso torna o fluxo de controle mais previsível, pois erros são valores comuns que podem ser retornados e manipulados diretamente.

## Definindo um error

Como **error** é uma interface podemos criar structs que a implementam.

```go
package main

import (
	"fmt"
)

type MyError struct {
	Message string
	Code    int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("error: %s, code: %d", e.Message, e.Code)
}

func ErrorTest(msg string, code int) error {
	return &MyError{Message: msg, Code: code}
}

func main() {
	err := ErrorTest("testando erro", -1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Jamais irei rodar :/")
}
```

Saída:
```txt
error: testando erro, code: -1
```

## Retornando um error

O Go segue uma abordagem idiomática para lidar com erros, que inclui a seguinte prática:

* Erros devem ser tratados logo após sua ocorrência:

    A maneira mais comum de lidar com erros em Go é verificá-los imediatamente após a chamada da função que pode falhar, como mostrado nos exemplos.

    ```go
    err := ErrorTest("testando erro", -1)
	if err != nil {
		fmt.Println(err)
		return
	}
    ```

    **OBS:** É extramamente recomendado não ignorar os erros.

* Erro como último retorno:

    A convenção no Go é retornar o erro como o último valor de retorno, depois do valor de sucesso (se houver). Isso facilita a leitura do código.

    ```go
    func ErrorTest(msg string, code int) (bool, error) {
	    return true, &MyError{Message: msg, Code: code}
    }

    func main() {
        out, err := ErrorTest("testando erro", -1)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("Jamais irei rodar :/")
    }
    ```

## Outras Abordagens

Existe várias formas de se lidar com error, seguem alguns exemplos:

**panic**

Em algumas situações críticas (erros irrecuperáveis), você pode usar **panic**. Isso interrompe a execução do programa imediatamente, gerando um stack trace. Entretanto, panic deve ser usado com parcimônia e apenas em situações realmente excepcionais.

```go
package main

import "fmt"

func main() {
    panic("Algo deu muito errado")
    fmt.Println("Isso nunca será impresso")
}
```

Saída:
```txt
panic: Algo deu muito errado

goroutine 1 [running]:
main.main()
        /home/rafaeldajuda/projects-go/tratamento-de-erro/main.go:30 +0x25
exit status 2
```

**recover**

Para recuperar de um panic, você pode usar a função **recover()** dentro de uma função **defer**. Isso permite capturar o erro e impedir que o programa seja encerrado abruptamente.

```go
package main

import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado de", r)
        }
    }()
    
    panic("Algo deu muito errado")
    fmt.Println("Isso nunca será impresso")
}
```

Saída:
```go
Recuperado de Algo deu muito errado
```

**fmt.Errorf**

A função **fmt.Errorf** uma nova mensagem de erro formatada.

```go
package main

import (
    "fmt"
)

func main() {
    err := fmt.Errorf("erro ao processar o arquivo %s", "data.txt")
    fmt.Println(err)
}
```

**errors.New**

A função **fmt.New** permite criar um novo erro com uma mesagem estática.

```go
package main

import (
    "errors"
    "fmt"
)

func main() {
    err := errors.New("ocorreu um erro")
    fmt.Println(err)
}
```

**errors.Is**

A função **errors.Is** em Go é usada para comparar erros, verificando se um erro específico é igual ou contém outro erro, mesmo que ele tenha sido "encapsulado" ou "aninhado". Isso é útil quando erros são envolvidos por outros erros, permitindo uma comparação precisa.

```go
package main

import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("recurso não encontrado")

func main() {
    err := fmt.Errorf("erro envolvido: %w", ErrNotFound)

    if errors.Is(err, ErrNotFound) {
        fmt.Println("Erro é ErrNotFound")
    }
}
```

**errors.As**

A função **errors.As** em Go é usada para verificar se um erro pode ser convertido para um tipo de erro específico e, se for possível, faz essa conversão. Ela é útil quando você precisa tratar erros personalizados ou tipos específicos de erros.

```go
package main

import (
    "errors"
    "fmt"
)

type MyError struct {
    Message string
}

func (e *MyError) Error() string {
    return e.Message
}

func main() {
    err := &MyError{Message: "ocorreu um erro personalizado"}
    var myErr *MyError

    if errors.As(err, &myErr) {
        fmt.Println("Erro é do tipo MyError:", myErr.Message)
    }
}
```

# Referências

- ChatGPT
- https://aprendagolang.com.br/dicas-e-praticas-sobre-error-handling-em-go/
