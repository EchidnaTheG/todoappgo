package utils

import (
"bufio"
 "os"
 "log"
 "strings"
 "strconv"
)



func GetFromDB(db *os.File) []Task{

	DBscanner := bufio.NewScanner(db)
	var toDoList []Task

	for DBscanner.Scan(){
		line := DBscanner.Text()
		if line == ""{
			continue
		}
		fields := strings.Split(line, ",")
		id, err := strconv.Atoi(fields[0])
		if err != nil{
			log.Fatal("Error at Getting From DB: ", err, "\nAre you sure you are using the DB rules?")
		}
		TextTask := fields[1]
		PriorityLevel, err := strconv.Atoi(fields[2])
		if err != nil{
			log.Fatal("Error at Getting From DB: ", err)
		}
		newTask := Task{id,TextTask,PriorityLevel}
		toDoList = append(toDoList,newTask )
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