package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var FILE_NAME string = "tasks.json"

func main() {

	data, err := os.ReadFile(FILE_NAME)
	if err != nil {
		if os.IsNotExist(err) {
			os.Create("tasks.json")
		} else {
			fmt.Println("Something unexpected happened.")
			return
		}
	}

	var tasks []Task

	if len(data) > 0 {
		err := json.Unmarshal(data, &tasks)
		if err != nil {
			fmt.Println("Something wrong happened when reading the json")
			return
		}
	}

	var choice string
	reader := bufio.NewReader(os.Stdin)

	for choice != "q" {
		fmt.Println("Choices:")
		fmt.Println("q: quit")
		fmt.Println("a: add task")
		fmt.Println("p: print tasks")
		fmt.Println("m: mark a task as complete")

		var input string
		input, _ = reader.ReadString('\n')
		choice = strings.TrimSpace(input)

		switch choice {
		case "a":
			tasks = handleAddTask(tasks, reader)
		case "p":
			printTasks(tasks)
		case "m":
			tasks = handleCompleteTask(tasks, reader)
		}

	}

}

func handleAddTask(tasks []Task, reader *bufio.Reader) []Task {
	fmt.Println("Enter the task you want to add: ")
	input, _ := reader.ReadString('\n')
	title := strings.TrimSpace(input)

	var id int

	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].Id + 1
	} else {
		id = 0
	}

	task := Task{Title: title, Completed: false, Id: id}

	tasks = append(tasks, task)

	writeTasksToFile(tasks, FILE_NAME)

	return tasks

}

func printTasks(tasks []Task) {
	for i := range tasks {
		var status string

		if tasks[i].Completed {
			status = "✅"
		} else {
			status = "❌"
		}
		fmt.Println("\n", "id: ", i, "\n", "Titolo: ", tasks[i].Title, "\n", "Status: ", status)
	}
}

func handleCompleteTask(tasks []Task, reader *bufio.Reader) []Task {
	fmt.Println("Enter the id of the task you want to mark as complete")
	printTasks(tasks)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	index, _ := strconv.Atoi(input)
	if 0 <= index && index < len(tasks) {
		tasks[index].Completed = true
	}
	writeTasksToFile(tasks, FILE_NAME)
	return tasks
}

func writeTasksToFile(tasks []Task, filename string) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
