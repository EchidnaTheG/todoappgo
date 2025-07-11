package utils

import (
	"fmt"
	"bufio"
	 "os"
	 "log"
	 "strconv"
	 "strings"
)


func welcomeMessage(){
	fmt.Println("Welcome To Your Task Management System!")
	fmt.Println("Please input 'Tasks' to see any current Tasks")
	fmt.Println("Please input 'Add' to add a new Task")
	fmt.Println("Please input 'Remove' to remove a completed Task")
	fmt.Println("Please input 'Exit()' to exit the program")
}


func tasks(toDoList *[]string){
	if len(*toDoList) == 0 {
				fmt.Println("No Tasks Pending!")
			
			}
			fmt.Println("Tasks To do:")
			for i, task := range *toDoList {
				fmt.Printf("  %v. %v\n", i+1, task)
			}
}


func add(toDoList *[]string, scanner *bufio.Scanner, db *os.File) *os.File {
	fmt.Println("What is your new task?")
	scanner.Scan()
	*toDoList = append(*toDoList, scanner.Text())
	db, err := os.Create(db.Name())
	if err != nil{
			log.Fatal("Error, Data Deleted, Please Verify ", err)
			}
	writer := bufio.NewWriter(db)
			for _, task := range *toDoList{
				_,err := writer.WriteString(task+"\n")
				if err != nil {
					log.Fatal("Error Writing: ",err)
				}
			}
			err = writer.Flush()
			if err != nil {
					log.Fatal("Error flushing: ",err)
				}
			fmt.Println("Task Added!")
			return db
}

func remove(toDoList *[]string, scanner *bufio.Scanner, db *os.File) *os.File {
		if len(*toDoList) == 0 {
			fmt.Println("No Tasks To Remove!")
			return db
		}

		
		scanCondition := true
		var index int
		for scanCondition{
			fmt.Println("Input Task Number to Remove")
			scanner.Scan()
			var err error
			index, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Printf("Error: %v\nVerify Index Number Carefully\n", err)
				continue
			}
			index -= 1

			if index > len(*toDoList)-1 || index < 0 {
				fmt.Println("Invalid Index Number")	
				continue
			}
			scanCondition = false
		}


		*toDoList = append((*toDoList)[:index], (*toDoList)[index+1:]...)
		db, err := os.Create(db.Name())
		writer := bufio.NewWriter(db)
		if err != nil{
			log.Fatal("Error In deleting Data ", err)
		}
		for _, task := range *toDoList{
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
			return db

}


func InnerLogic(db *os.File, toDoList *[]string) {
	
	scanner := bufio.NewScanner(os.Stdin)

	welcomeMessage()

	isOn := true

	for isOn {
		fmt.Println("\nInput New Command:")
		scanner.Scan()

		command := strings.ToLower(scanner.Text())
		fmt.Printf("You said: '%v'\n", command)


		switch command {

		case "tasks":
		tasks(toDoList)


		case "add":
			db=add(toDoList,scanner,db)


		case "remove":
			db = remove(toDoList, scanner, db)


		case "exit()":
			fmt.Println("Goodbye...")
			isOn = false
			continue


		default:
			fmt.Println("Unknown Command")
		}


	}
  }	

