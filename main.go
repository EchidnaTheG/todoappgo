package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func getFromDB(db *os.File) []string{
	DBscanner := bufio.NewScanner(db)
	
	var toDoList []string
	for DBscanner.Scan(){
		task := DBscanner.Text()
		toDoList = append(toDoList,task )
	}
	return toDoList
}

func VerifyDB() *os.File{
	db_name :="db.txt"
	if _, err := os.Stat(db_name); err != nil{
		db, err := os.Create(db_name)
		
		if err != nil{
			log.Fatal("Error in ",err)
		}
		return db
		
	}

	db, err := os.OpenFile(db_name,os.O_RDWR, 0644)
	if err != nil{
		log.Fatal("Error, Could Not OpenFile Line 108: ", err)
	}
	return db
	

	}

func InnerLogic(db *os.File, toDoList []string) {
	
	scanner := bufio.NewScanner(os.Stdin)

	welcomeMessage()
	writer := bufio.NewWriter(db)


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
			db, err := os.Create(db.Name())
			
			if err != nil{
				log.Fatal("Error, Data Deleted, Please Verify ", err)
			}
			writer = bufio.NewWriter(db)
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
			db, err = os.Create(db.Name())
			writer = bufio.NewWriter(db)
			if err != nil{
				log.Fatal("Error In deleting Data ", err)
			}

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
	db := VerifyDB()
	toDoList := getFromDB(db)
	InnerLogic(db, toDoList)
	defer db.Close()
}
