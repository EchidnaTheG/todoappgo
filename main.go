package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var toDoList []string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome To Your Task Management System!")
	fmt.Println("Please input 'Tasks' to see any current Tasks")
	fmt.Println("Please input 'Add' to add a new Task")
	fmt.Println("Please input 'Remove' to remove a completed Task")
	fmt.Println("Please input 'Exit()' to exit the program")
	isOn := true
	for isOn {
		fmt.Println("\nInput New Command:")
		scanner.Scan()
		command := strings.ToLower(scanner.Text())
		fmt.Printf("You said: '%v'\n", command)

		switch command {
		case "tasks":
			if len(toDoList) == 0 {
				fmt.Println("No Tasks Pending!")
				continue
			}
			fmt.Println("Tasks To do:")
			for i, task := range toDoList {
				fmt.Printf("  %v. %v\n", i+1, task)
			}
		case "add":
			fmt.Println("What is your new task?")
			scanner.Scan()
			toDoList = append(toDoList, scanner.Text())
			fmt.Println("Task Added!")
		case "remove":
			if len(toDoList) == 0 {
				fmt.Println("No Tasks To Remove!")
				continue
			}
			fmt.Println("Input Task Number to Remove")
			scanner.Scan()
			index, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Printf("Error: %v\nVerify Index Number Carefully\n", err)
				continue
			}
			index -= 1
			if index > len(toDoList)-1 || index < 0 {
				fmt.Println("Invalid Index Number")
				continue
			}
			toDoList = append(toDoList[:index], toDoList[index+1:]...)
			fmt.Println("Task Eliminated!")
		case "exit()":
			fmt.Println("Goodbye...")
			isOn = false
			continue
		default:
			fmt.Println("Unknown Command")
		}

	}
}
