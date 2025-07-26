package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	var choice string
	var tasks []string = make([]string, 0)

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
			tasks = handleAddTask(tasks)
		case "p":
			printTasks(tasks)
		case "m":
			tasks = handleCompleteTask(tasks)
		}

	}

}

func handleAddTask(tasks []string) []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the task you want to add: ")
	input, _ := reader.ReadString('\n')
	task := strings.TrimSpace(input)

	tasks = append(tasks, task)

	return tasks
}

func printTasks(tasks []string) {
	for i := range tasks {
		fmt.Println(i, tasks[i])
	}
}

func handleCompleteTask(tasks []string) []string {
	fmt.Println("Enter the number of the task you want to mark as complete")
	printTasks(tasks)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	index, _ := strconv.Atoi(input)
	if 0 <= index && index < len(tasks) {
		tasks = append(tasks[:index], tasks[index+1:]...)
	}
	return tasks
}
