package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)






func InnerLogic(db *os.File) {
	scanner := bufio.NewScanner(db)
	writer := bufio.NewWriter(db)
	var toDoList []string
	for scanner.Scan(){
		task := scanner.Text()
		toDoList = append(toDoList,task )
	}
	
	scanner = bufio.NewScanner(os.Stdin)



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
			os.Create(db.Name())
			for _, task := range toDoList{
				_,err := writer.WriteString(task+"\n")
				if err != nil {
					log.Fatal("Error Writing: ",err)
				}
			}
			err := writer.Flush()
			if err != nil {
					log.Fatal("Error flushing: ",err)
				}
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
			os.Create(db.Name())

			for _, task := range toDoList{
				_,err := writer.WriteString(task+"\n")
				if err != nil {
					log.Fatal("Error Writing: ",err)
				}
			}
			err = writer.Flush()
			if err != nil {
					log.Fatal("Error flushing: ",err)
				}
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




func main() {
	db_name :="db.txt"

	if _, err := os.Stat(db_name); err != nil{
		db, err := os.Create(db_name)
		defer db.Close()
		if err != nil{
			log.Fatal("Error in ",err)
		}
		InnerLogic(db)
	}

	db, err := os.OpenFile(db_name,os.O_RDWR, 0644)
	defer db.Close()
	if err != nil{
		log.Fatal("Error, Could Not OpenFile Line 108: ", err)
	}
	InnerLogic(db)


}
