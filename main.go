package main

import (
	"database/sql"
	"fmt"

	"github.com/RichSvK/Stock_Holder_Composition_Go/utility"
)

func main() {
	var db *sql.DB = nil
	var enterBuffer string = ""
	for db == nil {
		db = utility.LoginMenu()
		fmt.Print("Press [Enter] to continue...")
		fmt.Scanln(&enterBuffer)
		utility.ClearScreen()
	}
	// Close database when the main function end
	defer db.Close()

	var choice int = 0
	for choice != 3 {
		choice = utility.MainMenu()
		switch choice {
		case 1:
			utility.InsertMenu(db)
		case 2:
			utility.ExportMenu(db)
		default:
			fmt.Println("Program finished")
			return
		}
		fmt.Print("Press [Enter] to continue...")
		fmt.Scanln(&enterBuffer)
		utility.ClearScreen()
	}
}
