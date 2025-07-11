package utils

import (
"bufio"
 "os"
 "log"

)



func GetFromDB(db *os.File) []string{
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