package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var tasks map[string]Task
var status map[bool]string

type Task struct {
	ID        string
	Title     string
	Completed bool
}

func main() {

	// tasks list
	tasks = make(map[string]Task)
	status = make(map[bool]string)

	// task status
	status[false] = "[pendente]"
	status[true] = "[concluída]"

	// loop principal
	loop := true
	for loop {
		// mensagem de entrada
		fmt.Println("Comandos disponíveis: add, list, done, delete, exit")

		reader := bufio.NewReader(os.Stdin)

		// entrada do comando
		comando, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Erro na leitura do comando: %s\n", err.Error())
			return
		}
		// fmt.Printf("comando: %#v\n", comando)
		cmd, task := getInfo(strings.TrimSpace(comando))

		switch strings.ToLower(cmd) {
		case "add":
			addTask(task)
		case "list":
			listTasks()
		case "done":
			doneTask(task)
		case "delete":
			delete(tasks, task)
			fmt.Printf("Tarefa %s removida\n", task)
		case "exit":
			fmt.Println("Encerrando programa...")
			loop = false
		default:
			fmt.Println("Comando inválido")
		}
	}

}

func getInfo(command string) (cmd string, task string) {
	index := strings.Index(command, " ")
	cmd = command
	if index > 1 {
		cmd = command[:index]
		task = command[(index + 1):]
	}

	return
}

func genID() string {
	for {
		id := strconv.FormatInt(rand.Int63n(99999), 10)
		if tasks[id] == (Task{}) {
			return id
		}
	}
}

func addTask(task string) {
	id := genID()
	item := Task{ID: id, Title: task, Completed: false}
	tasks[id] = item
	fmt.Printf("Tarefa adicionada com ID %s\n", id)
}

func listTasks() {
	for _, item := range tasks {
		fmt.Printf("%s - %s %s\n", item.ID, item.Title, status[item.Completed])
	}
}

func doneTask(task string) {
	if tasks[task] != (Task{}) { // nesse caso task vai ser o ID
		item := tasks[task]
		item.Completed = true
		tasks[task] = item
		fmt.Printf("Tarefa %s marcada como concluída\n", task)
	} else {
		fmt.Printf("Tarefa %s não encontrada\n", task)
	}
}
