package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	ID        int
	Title     string
	Author    string
	Available bool
}

var library []Book

const (
	BORROW_BOOK    = "Emprestado"
	AVAILABLE_BOOK = "Disponível"
)

// start library
func init() {
	library = append(library, Book{ID: 1, Title: "O Enigma da Meia-Noite", Author: "Lúcia M. Ferreira", Available: true})
	library = append(library, Book{ID: 2, Title: "Caminhos do Vento", Author: "João P. Andrade", Available: false})
	library = append(library, Book{ID: 3, Title: "Sombras do Passado", Author: "Renata T. Lima", Available: true})
	library = append(library, Book{ID: 4, Title: "O Código Esquecido", Author: "Carlos D. Mendes", Available: true})
	library = append(library, Book{ID: 5, Title: "Além do Horizonte", Author: "Fernanda S. Ribeiro", Available: false})
	library = append(library, Book{ID: 6, Title: "A Última Profecia", Author: "Marcos L. Oliveira", Available: true})
	library = append(library, Book{ID: 7, Title: "Reflexos da Alma", Author: "Patrícia G. Souza", Available: true})
	library = append(library, Book{ID: 8, Title: "Diário de um Sonhador", Author: "Bruno R. Teixeira", Available: false})
	library = append(library, Book{ID: 9, Title: "Ecos do Silêncio", Author: "Adriana V. Cunha", Available: true})
	library = append(library, Book{ID: 10, Title: "O Guardião das Estrelas", Author: "Thiago M. Santos", Available: false})
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	loop := true
	for loop {
		fmt.Println("Commands: add <titulo>;<autor>, list, borrow <id>, return <id>, remove <id>, exit")

		option, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error get option:", err.Error())
			return
		}
		cmd, value, err := getCommand(strings.TrimSpace(option))
		if err != nil {
			cmd = err.Error()
		}

		switch strings.ToLower(cmd) {
		case "add":
			addBook(value)
		case "list":
			listBooks()
		case "borrow":
			borrowBook()
		case "return":
			returnBook()
		case "remove":
			removeBook()
		case "exit":
			fmt.Println("Encerrando programa...")
			loop = false
		default:
			fmt.Println("Comando inválido!")
		}
	}
}

func getCommand(command string) (string, string, error) {
	if strings.Count(command, " ") > 1 {
		return "", "", fmt.Errorf("comando inválido")
	}
	index := strings.Index(command, " ")
	cmd := command
	if index > 1 {
		cmd = command[:index]
		value := command[(index + 1):]
		return cmd, value, nil
	}

	return cmd, "", nil
}

// library funcs
func addBook(value string) {
	fmt.Println("addBook")
}
func listBooks() {
	fmt.Println("listBooks")
}
func borrowBook() {
	fmt.Println("borrowBook")
}
func returnBook() {
	fmt.Println("returnBook")
}
func removeBook() {
	fmt.Println("removeBook")
}
