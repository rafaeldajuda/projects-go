package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// lendo a entrada de texto
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println("Lendo:", text)

	// tratando texto - deixando minúsculo - tirando pontuações
	var punctuationList []string = []string{".", ",", ";", ":", "?", "!", "...", "-", "–", "\"", "'", "(", ")", "[", "]", "{", "}"}
	lowerText := strings.ToLower(text)
	for _, value := range punctuationList {
		lowerText = strings.ReplaceAll(lowerText, value, "")
	}

	// contando palavras únicas
	var wordMap = make(map[string]bool)
	wordList := strings.Fields(lowerText)
	for _, value := range wordList {
		if !wordMap[value] { // se não estivar na lista, adiciona-la
			wordMap[value] = true
		}
	}

	// saída
	fmt.Print("Palavras únicas:")
	for key := range wordMap {
		fmt.Print(" " + key)
	}
	fmt.Println()
	fmt.Println(len(wordMap))

}
