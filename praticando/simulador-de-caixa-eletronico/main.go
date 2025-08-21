package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Extratos struct {
	Registro []Registro
}

type Registro struct {
	Operacao string
	Valor    float64
}

var extratos Extratos

const (
	OPERACAO_DEPOSITO = "Deposito"
	OPERACAO_SAQUE    = "Saque"
)

func init() {
	// Instânciando extratos com saldo de R$1000
	extratos.Registro = append(extratos.Registro, Registro{Operacao: OPERACAO_DEPOSITO, Valor: 1000})
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	loop := true
	for loop {
		fmt.Println("Comandos: saldo, depositar <valor>, sacar <valor>, extrato, sair")

		comando, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("erro leitura comando:", err.Error())
			return
		}
		cmd, valor, _ := GetComando(strings.TrimSpace(comando))

		switch strings.ToLower(cmd) {
		case "saldo":
			saldo := GetSaldo()
			fmt.Printf("Saldo atual: R$%.2f\n", saldo)
		case "depositar":
			valido := Depositar(valor)
			if !valido {
				fmt.Println("Comando inválido!")
				continue
			}
			fmt.Printf("Depósito de R$%.2f realizado com sucesso\n", valor)
		case "sacar":
			valido := Sacar(valor)
			if !valido {
				fmt.Println("Comando inválido!")
				continue
			}
			fmt.Printf("Saque de R$%.2f realizado com sucesso\n", valor)
		case "extrato":
			Extrato()
		case "sair":
			fmt.Println("Encerrando programa...")
			loop = false
		default:
			fmt.Println("Comando inválido!")
		}
	}

}

func GetComando(comando string) (string, float64, error) {
	if strings.Count(comando, " ") > 1 {
		return "", 0, fmt.Errorf("comando inválido")
	}
	index := strings.Index(comando, " ")
	cmd := comando
	if index > 1 {
		cmd = comando[:index]
		valor, err := strconv.ParseFloat(comando[(index+1):], 64)
		if err != nil {
			return "", 0, fmt.Errorf("valor inválido: %v", err.Error())
		}
		return cmd, valor, nil
	}

	return cmd, 0, nil
}

func Depositar(valor float64) bool {
	if valor <= 0 {
		return false
	}
	extratos.Registro = append(extratos.Registro, Registro{Operacao: OPERACAO_DEPOSITO, Valor: valor})
	return true
}

func Sacar(valor float64) bool {
	if valor > GetSaldo() || valor <= 0 {
		return false
	}
	extratos.Registro = append(extratos.Registro, Registro{Operacao: OPERACAO_SAQUE, Valor: valor})
	return true
}
func Extrato() {
	for _, item := range extratos.Registro {
		fmt.Printf("[%s: %.2f]\n", item.Operacao, item.Valor)
	}
}

func GetSaldo() float64 {
	saldo := 0.
	for _, item := range extratos.Registro {
		if item.Operacao == OPERACAO_DEPOSITO {
			saldo += item.Valor
		} else if item.Operacao == OPERACAO_SAQUE {
			saldo -= item.Valor
		}
	}

	return saldo
}
