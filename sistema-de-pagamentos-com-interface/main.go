package main

import "fmt"

// Definição da interface PaymentMethod com a função Pay
type PaymentMethod interface {
	Pay(amount float64) string
}

// Estrutura CreditCard com o campo CardNumber
type CreditCard struct {
	CardNumber string
}

// Implementação da função Pay que retorna o pagamento e o número do cartão
func (c CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Pagamento realizado com sucesso: R$%.2f. Número do cartão: %s", amount, c.CardNumber)
}

// Estrutura Paypal com o campo Email
type Paypal struct {
	Email string
}

// Implementação da função Pay que retorna o pagamento e o email do usuário
func (p Paypal) Pay(amount float64) string {
	return fmt.Sprintf("Pagamento realizado com sucesso: R$%.2f. Conta: %s", amount, p.Email)
}

// Estrutura Bitcoin com o campo WalletAddress
type Bitcoin struct {
	WalletAddress string
}

// Implementação da função Pay que retorna o pagamento e a carteira do usuário
func (b Bitcoin) Pay(amount float64) string {
	return fmt.Sprintf("Pagamento realizado com sucesso: R$%.2f. Carteira: %s", amount, b.WalletAddress)
}

// Função que realiaza vários pagamentos a partir de uma lista de PaymentMethod
func Payments(methods []PaymentMethod, amount float64) {
	for _, method := range methods {
		fmt.Println(method.Pay(amount))
	}
}

func main() {
	// Lista de métodos de pagamentos
	methods := []PaymentMethod{CreditCard{CardNumber: "1234-2345-3456-4567"},
		Paypal{Email: "rafael@gmail.com"},
		Bitcoin{WalletAddress: "rafa123Adress"},
	}

	// Chamando a função Payments, passando a lista de pagamentos e o valor a ser pago
	Payments(methods, 9.99)
}
