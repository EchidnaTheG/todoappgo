package main

import (
	"github.com/EchidnaTheG/mytodoApp/utils"
)







func main() {
	db := utils.VerifyDB()
	defer db.Close()
	toDoList := utils.GetFromDB(db)
	utils.InnerLogic(db, &toDoList)
	
}
