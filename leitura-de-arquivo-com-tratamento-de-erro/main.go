package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// Tenta abrir um arquivo cujo nome está vazio, então vai dar erro.
	file, err := os.Open("")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err.Error())
		return
	}
	defer file.Close() // Garante que o arquivo será fechado no final da função.

	// Lê todo o conteúdo do arquivo.
	text, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Erro na leitura do arquivo:", err.Error())
		return
	}

	// Converte o conteúdo lido para string e imprime no console.
	fmt.Println(string(text))

}
