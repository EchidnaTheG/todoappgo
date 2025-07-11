package main

import (
	"github.com/EchidnaTheG/mytodoApp/utils"
)







func main() {
	db := utils.VerifyDB()
	toDoList := utils.GetFromDB(db)
	utils.InnerLogic(db, &toDoList)
	defer db.Close()
}
