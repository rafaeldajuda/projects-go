package main

import "fmt"

// Define um tipo de erro personalizado para divisão por zero
type ErrorDivideByZero struct {
	divisor float64 // Armazena o valor do divisor que causou o erro
}

// Implementa o método Error, que retorna uma mensagem de erro
func (e *ErrorDivideByZero) Error() string {
	return "error: divide by zero (divisor)"
}

// Função Divide realiza a divisão entre dois números float64
// Retorna um erro personalizado se o divisor for zero
func Divide(a float64, b float64) (result float64, err error) {
	if b == 0 {
		return result, &ErrorDivideByZero{divisor: b}
	}
	return a / b, nil
}

func main() {

	// operação com erro
	a, b := 3., 0.

	_, err := Divide(a, b)
	if err != nil {
		fmt.Println(err.Error())
	}

	// operação com sucesso
	a, b = 3., 2.

	result, err := Divide(a, b)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("result:", result)

}
